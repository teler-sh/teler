package parsers

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type options struct {
	Active   bool     `yaml:"active"`
	Excludes []string `yaml:"excludes"`
}

type slack struct {
	Text    string
	URL     string
	Token   string `yaml:"token"`
	Color   string `yaml:"color"`
	Channel string `yaml:"channel"`
}

type telegram struct {
	Text      string
	URL       string
	Token     string `yaml:"token"`
	ChatID    string `yaml:"chat_id"`
	Silent    bool   `yaml:"silent"`
	ParseMode string `yaml:"parse_mode"`
}

// Config default structure for config
type Config struct {
	Configs struct {
		Format string `yaml:"format" validate:"nonzero"`

		Rules struct {
			Threat options `yaml:"threat" validate:"nonzero"`
			Filter options `yaml:"filter" validate:"nonzero"`
		} `yaml:"rules" validate:"nonzero"`

		Notification struct {
			Active   bool   `yaml:"active"`
			Provider string `yaml:"provider"`
		} `yaml:"notification" validate:"nonzero"`
	} `yaml:"configs" validate:"nonzero"`

	Notifications struct {
		Slack    slack    `yaml:"slack"`
		Telegram telegram `yaml:"telegram"`
	} `yaml:"notifications"`
}

// GetConfig will parse the config file
func GetConfig(f string) (*Config, error) {
	config := &Config{}
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
