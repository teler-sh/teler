package parsers

type options struct {
	Excludes   []string `yaml:"excludes"`
	Whitelists []string `yaml:"whitelists"`
}
