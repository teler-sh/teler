package parsers

import (
	"os"

	"gopkg.in/yaml.v2"
)

// GetConfig will parse the config file
func GetConfig(f string) (*Configs, error) {
	config := &Configs{}
	file, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	err = GetYaml(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetYaml file configuration
func GetYaml(f []byte, s interface{}) error {
	y := yaml.Unmarshal(f, s)
	return y
}
