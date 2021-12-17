package parsers

type prometheus struct {
	Active   bool   `yaml:"active"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Endpoint string `yaml:"endpoint"`
}
