package requests

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/briandowns/spinner"
	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/pkg/cache"
	"github.com/kitabisa/teler/pkg/errors"
	"github.com/kitabisa/teler/resource"
)

var (
	rsrc    *resource.Resources
	content []byte
	errCon  error
)

// Resources is to getting all available resources
func Resources(options *common.Options) {
	rsrc = resource.Get()
	getRules(options)
}

func getRules(options *common.Options) {
	var spin = spinner.New(spinner.CharSets[11], 90*time.Millisecond, spinner.WithWriter(os.Stderr))
	var exclude bool

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
		fname := rsrc.Threat[i].Filename
		cat := rsrc.Threat[i].Category

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
			content, errCon = os.ReadFile(filepath.Join(cache.Path, fname))
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

			content, errCon = io.ReadAll(res.Body)
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
