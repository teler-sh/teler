package teler

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
	"github.com/logrusorgru/aurora"
	"github.com/satyrius/gonx"
	"github.com/valyala/fastjson"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/configs"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/matchers"
)

// Analyze logs from threat resources
func Analyze(options *common.Options, logs *gonx.Entry) {
	var match bool
	var threatCat, threatElm string

	log := make(map[string]string)
	resource := configs.Get()

	fields := reflect.ValueOf(logs).Elem().FieldByName("fields")
	for _, field := range fields.MapKeys() {
		log[field.String()] = fields.MapIndex(field).String()
	}

	go func() {
		for i := 0; i < len(resource.Threat); i++ {
			threat := reflect.ValueOf(&resource.Threat[i]).Elem()
			cat := threat.FieldByName("Category").String()
			con := threat.FieldByName("Content").String()
			exc := threat.FieldByName("Exclude").Bool()

			if exc {
				continue
			}

			switch cat {
			case "Common Web Attack":
				req, _ := url.Parse(log["request_uri"])
				query := req.Query()
				if len(query) > 0 {
					for _, q := range query {
						fil, _ := fastjson.Parse(con)
						dec, _ := url.QueryUnescape(strings.Join(q, ""))
						cwa := fil.GetArray("filters")
						for _, v := range cwa {
							match = matchers.IsMatch(string(v.GetStringBytes("rule")), regexp.QuoteMeta(dec))
							threatCat = cat + ": " + string(v.GetStringBytes("description"))
							threatElm = log["request_uri"]

							if match {
								break
							}
						}
					}
				}
			case "Bad Crawler":
				for _, pat := range strings.Split(con, "\n") {
					match = matchers.IsMatch(pat, log["http_user_agent"])
					if match {
						break
					}
				}
				threatCat = cat
				threatElm = log["http_user_agent"]
			case "Bad IP Address":
				ip := "(?m)^" + log["remote_addr"]
				match = matchers.IsMatch(ip, con)
				threatCat = cat
				threatElm = log["remote_addr"]
			case "Bad Referrer":
				if log["http_referer"] == "-" {
					break
				}

				req, _ := url.Parse(log["http_referer"])
				ref := "(?m)^"
				ref += req.Path
				if req.Host != "" {
					ref += req.Host
				}

				match = matchers.IsMatch(ref, con)
				threatCat = cat
				threatElm = log["http_referer"]
			case "Directory Bruteforce":
				req, _ := url.Parse(log["request_uri"])
				if matchers.IsMatch("^(2|3)[0-9]{2}", log["status"]) {
					break
				}

				if req.Path != "/" {
					match = matchers.IsMatch(trimFirst(req.Path), con)
					threatCat = cat
					threatElm = log["request_uri"]
				}
			}

			if match {
				out := fmt.Sprintf("[%s] [%s] %s", aurora.Cyan(log["time_local"]), aurora.Yellow(threatCat), aurora.Red(threatElm))
				if options.Output != "" {
					if _, write := options.OutFile.WriteString(fmt.Sprintf("%s\n", stripansi.Strip(out))); write != nil {
						errors.Show(write.Error())
					}
				}
				fmt.Println(out)
				// if options.Configs.Alert.Active {
				// 	sendAlert(options, log)
				// }
			}
		}
	}()
}

func trimFirst(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
