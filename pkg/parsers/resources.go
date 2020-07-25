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
	getf, _ := ioutil.ReadFile(file)

	err := GetYaml(getf, rsrc)
	if err != nil {
		return nil, err
	}

	return rsrc, nil
}
