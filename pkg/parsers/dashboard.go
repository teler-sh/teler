package parsers

type dashboard struct {
	Active   bool   `yaml:"active"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Endpoint string `yaml:"endpoint"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
