package configs

type config []struct {
	Category string `yaml:"cat"`
	URL      string `yaml:"url"`
	Content  string
}

type Resources struct {
	Threat config `yaml:"threat"`
	Filter config `yaml:"filter"`
}
