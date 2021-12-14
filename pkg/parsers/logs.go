package parsers

type zinc struct {
	Active bool   `yaml:"active"`
	SSL    bool   `yaml:"ssl"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Index  string `yaml:"index"`
}

type file struct {
	Active bool   `yaml:"active"`
	JSON   bool   `yaml:"json"`
	Path   string `yaml:"path"`
}
