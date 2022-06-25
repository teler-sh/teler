package parsers

// Configs default structure for configurations
type Configs struct {
	Logformat string `yaml:"log_format" validate:"nonzero"`

	Rules struct {
		Cache  bool    `yaml:"cache"`
		Threat options `yaml:"threat" validate:"nonzero"`
	} `yaml:"rules" validate:"nonzero"`

	Dashboard dashboard `yaml:"dashboard" validate:"nonzero"`

	Metrics struct {
		Prometheus prometheus `yaml:"prometheus"`
	} `yaml:"metrics" validate:"nonzero"`

	Logs struct {
		File file `yaml:"file"`
		Zinc zinc `yaml:"zinc"`
	} `yaml:"logs" validate:"nonzero"`

	Alert struct {
		Active   bool   `yaml:"active"`
		Provider string `yaml:"provider"`
	} `yaml:"alert" validate:"nonzero"`

	Notifications struct {
		Slack    general  `yaml:"slack"`
		Telegram telegram `yaml:"telegram"`
		Discord  general  `yaml:"discord"`
	} `yaml:"notifications"`
}
