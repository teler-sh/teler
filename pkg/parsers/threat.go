package parsers

type customs struct {
	Name      string `yaml:"name"`
	Condition string `yaml:"condition"`
	Rules     []struct {
		Element  string `yaml:"element"`
		Pattern  string `yaml:"pattern"`
		Selector bool   `yaml:"selector"`
	} `yaml:"rules"`
}

type options struct {
	Excludes   []string  `yaml:"excludes"`
	Whitelists []string  `yaml:"whitelists"`
	Customs    []customs `yaml:"customs"`
}
