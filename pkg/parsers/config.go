package parsers

import "io/ioutil"

type options struct {
	Excludes   []string `yaml:"excludes"`
	Whitelists []string `yaml:"whitelists"`
}

type general struct {
	Token   string `yaml:"token"`
	Color   string `yaml:"color"`
	Channel string `yaml:"channel"`
}

type telegram struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chat_id"`
}

// Configs default structure for config
type Configs struct {
	Logformat string `yaml:"log_format" validate:"nonzero"`

	Rules struct {
		Cache  bool    `yaml:"cache"`
		Threat options `yaml:"threat" validate:"nonzero"`
	} `yaml:"rules" validate:"nonzero"`

	Prometheus struct {
		Active   bool   `yaml:"active"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"prometheus" validate:"nonzero"`

	Alert struct {
		Active   bool   `yaml:"active"`
		Provider string `yaml:"provider"`
	} `yaml:"alert" validate:"nonzero"`

	Notifications struct {
		Slack    general    `yaml:"slack"`
		Telegram telegram `yaml:"telegram"`
		Discord  general  `yaml:"discord"`
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
