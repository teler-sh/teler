package runner

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"

	"gopkg.in/validator.v2"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/matchers"
	"ktbs.dev/teler/pkg/parsers"
	"ktbs.dev/teler/pkg/utils"
)

func validate(options *common.Options) {
	if !options.Stdin {
		if options.Input == "" {
			errors.Exit(errors.ErrNoInputLog)
		}
	}

	if options.ConfigFile == "" {
		telerEnv := os.Getenv("TELER_CONFIG")
		if telerEnv == "" {
			errors.Exit(errors.ErrNoInputConfig)
		} else {
			options.ConfigFile = telerEnv
		}
	}

	config, errConfig := parsers.GetConfig(options.ConfigFile)
	if errConfig != nil {
		errors.Exit(errors.ErrParseConfig + errConfig.Error())
	}

	if config.Logs.File.Active {
		if config.Logs.File.Path == "" {
			errors.Exit(errors.ErrNoFilePath)
		}

		f, errOutput := os.OpenFile(config.Logs.File.Path,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if errOutput != nil {
			errors.Exit(errOutput.Error())
		}
		options.Output = f
	}

	// Validates log format
	matchers.IsLogformat(config.Logformat)
	options.Configs = config

	// Validates custom threat rules & notification parts configuration
	customs(options)
	notification(options)

	// Do Zinc health check, validate & set credentials
	if config.Logs.Zinc.Active {
		options.Configs.Logs.Zinc.Base64Auth = zinc(options)
	}

	if errVal := validator.Validate(options); errVal != nil {
		errors.Exit(errVal.Error())
	}
}

func customs(options *common.Options) {
	var err string

	cfg := options.Configs
	cat := make(map[string]bool)

	custom := cfg.Rules.Threat.Customs
	for i := 0; i < len(custom); i++ {
		cond := strings.ToLower(custom[i].Condition)
		matchers.IsCondition(cond)
		matchers.IsBlank(custom[i].Name, "Custom threat category")

		if cat[custom[i].Name] {
			err = strings.Replace(errors.ErrDupeCategory, ":category", custom[i].Name, -1)
			errors.Exit(err)
		}
		cat[custom[i].Name] = true

		rules := custom[i].Rules
		if len(rules) < 1 {
			err = strings.Replace(errors.ErrNoThreatRules, ":category", custom[i].Name, -1)
			errors.Exit(err)
		}

		for j := 0; j < len(rules); j++ {
			matchers.IsBlank(rules[j].Element, "Custom threat rules element")
			elm := fmt.Sprint("$", rules[j].Element)

			if !matchers.IsMatch(fmt.Sprint(`\`, elm), cfg.Logformat) {
				err = strings.Replace(errors.ErrNoElement, ":element", elm, -1)
				err = strings.Replace(err, ":category", custom[i].Name, -1)

				errors.Exit(err)
			}

			matchers.IsBlank(rules[j].Pattern, "Custom threat rules pattern")
		}
	}
}

func zinc(options *common.Options) string {
	var health, auth map[string]interface{}

	zinc := options.Configs.Logs.Zinc
	base := "http"
	if zinc.SSL {
		base += "s"
	}
	base += fmt.Sprint("://", zinc.Host, ":", zinc.Port)

	resp, err := http.Get(fmt.Sprint(base, "/healthz"))
	if err != nil {
		errors.Exit(fmt.Sprint(errors.ErrHealthZinc, ": ", err.Error()))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Exit(fmt.Sprint(errors.ErrHealthZinc, ": ", err.Error()))
	}

	if err = json.Unmarshal(body, &health); err != nil {
		errors.Exit(fmt.Sprint(errors.ErrHealthZinc, ": ", err.Error()))
	}

	if health["status"] != "ok" {
		errors.Exit(errors.ErrHealthZinc)
	}

	b64auth := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprint(zinc.Username, ":", zinc.Password),
	))
	data, _ := json.Marshal(map[string]string{
		"_id":           zinc.Username,
		"base64encoded": b64auth,
		"password":      zinc.Password,
	})

	resp, err = http.Post(
		fmt.Sprint(base, "/api/login"),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		errors.Exit(fmt.Sprint(errors.ErrAuthZinc, ": ", err.Error()))
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Exit(fmt.Sprint(errors.ErrAuthZinc, ": ", err.Error()))
	}

	if err = json.Unmarshal(body, &auth); err != nil {
		errors.Exit(fmt.Sprint(errors.ErrAuthZinc, ": ", err.Error()))
	}

	if auth["validated"] == false {
		errors.Exit(errors.ErrAuthZinc)
	}

	return b64auth
}

func notification(options *common.Options) {
	var useWebhook bool
	config := options.Configs

	if config.Alert.Active {
		provider := utils.Title(config.Alert.Provider)
		field := reflect.ValueOf(&config.Notifications).Elem().FieldByName(provider)

		switch provider {
		case "Slack", "Discord":
			if matchers.IsWebhook(provider, field.FieldByName("Webhook").String()) {
				useWebhook = true
			} else {
				matchers.IsChannel(field.FieldByName("Channel").String())
			}

			matchers.IsColor(field.FieldByName("Color").String())
		case "Telegram":
			matchers.IsChatID(field.FieldByName("ChatID").String())
		default:
			errors.Exit(strings.Replace(errors.ErrAlertProvider, ":platform", config.Alert.Provider, -1))
		}

		if !useWebhook {
			matchers.IsToken(field.FieldByName("Token").String())
		}
	}
}

func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}
	return true
}
