package requests

import (
	"io/ioutil"
	"net/http"
	"reflect"

	log "github.com/projectdiscovery/gologger"
	"ktbs.me/teler/common"
	"ktbs.me/teler/configs"
)

var resource *configs.Resources
var exclude bool

// Resources is to getting all available resources
func Resources(options *common.Options) {
	resource = configs.Get()
	getRules(options)
}

func getRules(options *common.Options) {
	client := Client()
	excludes := options.Configs.Rules.Threat.Excludes

	for i := 0; i < len(resource.Threat); i++ {
		exclude = false
		threat := reflect.ValueOf(&resource.Threat[i]).Elem()

		for x := 0; x < len(excludes); x++ {
			if excludes[x] == threat.FieldByName("Category").String() {
				exclude = true
			}
			threat.FieldByName("Exclude").SetBool(exclude)
		}

		if exclude {
			continue
		}

		log.Infof("Getting \"%s\" resource...\n", threat.FieldByName("Category").String())

		req, _ := http.NewRequest("GET", threat.FieldByName("URL").String(), nil)
		resp, _ := client.Do(req)

		body, _ := ioutil.ReadAll(resp.Body)
		threat.FieldByName("Content").SetString(string(body))
	}
}
