package requests

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/briandowns/spinner"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/cache"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/resource"
)

var (
	rsrc    *resource.Resources
	exclude bool
	content []byte
	errCon  error
	spin    *spinner.Spinner
)

func init() {
	spin = spinner.New(spinner.CharSets[11], 90*time.Millisecond, spinner.WithWriter(os.Stderr))
}

// Resources is to getting all available resources
func Resources(options *common.Options) {
	rsrc = resource.Get()
	getRules(options)
}

func getRules(options *common.Options) {
	client := Client()

	rules := options.Configs.Rules
	excludes := rules.Threat.Excludes
	isCached := rules.Cache

	if err := spin.Color("blue"); err != nil {
		errors.Exit(err.Error())
	}

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

		spin.Suffix = " Getting \"" + cat + "\" resource..."

		if cache.Check() && isCached {
			content, errCon = ioutil.ReadFile(filepath.Join(cache.Path, fname))
			if errCon != nil {
				cache.Purge()

				println()
				errors.Show("Fail to get local resources. Retry...")
				getRules(options)

				spin.Restart()
			}
		} else {
			req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/kitabisa/teler-resources/master/db/"+fname, nil)
			if err != nil {
				errors.Exit(err.Error())
			}

			res, err := client.Do(req)
			if err != nil {
				errors.Exit(err.Error())
			}

			content, errCon = ioutil.ReadAll(res.Body)
			if errCon != nil {
				errors.Exit(errCon.Error())
			}

			if isCached {
				file, err := os.Create(filepath.Join(cache.Path, fname))
				if err != nil {
					errors.Exit(err.Error())
				}

				if _, err = file.WriteString(string(content)); err != nil {
					errors.Exit(err.Error())
					file.Close()
				}
			}
		}

		threat.FieldByName("Content").SetString(string(content))
	}

	spin.Stop()

	if isCached {
		cache.Update()
	}
}
