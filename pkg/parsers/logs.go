package parsers

type zinc struct {
	Active     bool   `yaml:"active"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	SSL        bool   `yaml:"ssl"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Index      string `yaml:"index"`
	Base64Auth string
}

type file struct {
	Active bool   `yaml:"active"`
	JSON   bool   `yaml:"json"`
	Path   string `yaml:"path"`
}
