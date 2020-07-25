package parsers

import (
	"gopkg.in/yaml.v2"
)

func GetYaml(f []byte, s interface{}) error {
	y := yaml.Unmarshal(f, s)
	return y
}
