package requests

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/projectdiscovery/gologger"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/cache"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/resource"
)

var (
	rsrc    *resource.Resources
	exclude bool
	content []byte
)

// Resources is to getting all available resources
func Resources(options *common.Options) {
	rsrc = resource.Get()
	getRules(options)
}

func getRules(options *common.Options) {
	client := Client()
	excludes := options.Configs.Rules.Threat.Excludes

	for i := 0; i < len(rsrc.Threat); i++ {
		exclude = false
		threat := reflect.ValueOf(&rsrc.Threat[i]).Elem()
		fname := threat.FieldByName("Filename").String()
		cat := threat.FieldByName("Category").String()

		for x := 0; x < len(excludes); x++ {
			if excludes[x] == cat {
				exclude = true
			}
			threat.FieldByName("Exclude").SetBool(exclude)
		}

		if exclude {
			continue
		}

		gologger.Infof("Getting \"%s\" resource...\n", cat)

		if cache.IsCached() {
			content, _ = ioutil.ReadFile(filepath.Join(cache.Path, fname))
		} else {
			req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/kitabisa/teler-resources/master/db/"+fname, nil)
			if err != nil {
				errors.Exit(err.Error())
			}

			res, err := client.Do(req)
			if err != nil {
				errors.Exit(err.Error())
			}

			content, _ = ioutil.ReadAll(res.Body)

			file, err := os.Create(filepath.Join(cache.Path, fname))
			if err != nil {
				errors.Exit(err.Error())
			}

			if _, err = file.WriteString(string(content)); err != nil {
				errors.Exit(err.Error())
				file.Close()
			}
		}

		threat.FieldByName("Content").SetString(string(content))
	}

	cache.Update()
}
