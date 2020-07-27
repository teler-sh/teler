package parsers

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"

	"github.com/kitabisa/teler/configs"
)

// GetResources read resource file configuration
func GetResources() (*configs.Resources, error) {
	rsrc := &configs.Resources{}
	genv := os.Getenv("GOPATH")
	conf := reflect.TypeOf(configs.Resources{}).PkgPath()
	file := path.Join(genv, "src", conf, "resources.yaml")
	getf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	yaml := GetYaml(getf, rsrc)
	if yaml != nil {
		return nil, yaml
	}

	return rsrc, nil
}
