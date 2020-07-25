package configs

type config []struct {
	Cat string `yaml:"cat"`
	URL string `yaml:"url"`
}

type Resources struct {
	Threat config `yaml:"threat"`
	Filter config `yaml:"filter"`
}
