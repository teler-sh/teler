package parsers

import "io/ioutil"

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

// Configs default structure for config
type Configs struct {
	Format string `yaml:"format" validate:"nonzero"`

	Rules struct {
		Threat options `yaml:"threat" validate:"nonzero"`
		Filter options `yaml:"filter" validate:"nonzero"`
	} `yaml:"rules" validate:"nonzero"`

	Alert struct {
		Active   bool   `yaml:"active"`
		Provider string `yaml:"provider"`
	} `yaml:"notification" validate:"nonzero"`

	Notifications struct {
		Slack    slack    `yaml:"slack"`
		Telegram telegram `yaml:"telegram"`
	} `yaml:"notifications"`
}

// GetConfig will parse the config file
func GetConfig(f string) (*Configs, error) {
	config := &Configs{}
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	err = GetYaml(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
