package teler

import (
	"reflect"
	"strings"

	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/configs"
	"github.com/kitabisa/teler/pkg/matchers"
	"github.com/projectdiscovery/gologger"
	"github.com/satyrius/gonx"
)

var ex, m bool
var time_local string

// Analyze logs from threat resources
func Analyze(options *common.Options, logs *gonx.Entry) {
	resource := configs.Get()
	fields := reflect.ValueOf(logs).Elem().FieldByName("fields")
	for _, field := range fields.MapKeys() {
		log := fields.MapIndex(field).String()
		cat := field.String()
		if cat == "time_local" {
			time_local = log
		}

		for i := 0; i < len(resource.Threat); i++ {
			threat := reflect.ValueOf(&resource.Threat[i]).Elem()
			category := threat.FieldByName("Category").String()
			content := threat.FieldByName("Content").String()

			patterns := strings.Split(content, "\n")
			for _, pattern := range patterns {
				pattern = strings.TrimSpace(pattern)
				if pattern == "" {
					continue
				}

				switch category {
				case "Common Web Attack":
				case "Bad Crawler":
					switch {
					case cat == "http_user_agent":
						detect(options, category, pattern, log, time_local)
					}
				default:
					switch {
					case cat == "remote_addr":
						if category == "Bad IP Address" {
							detect(options, category, pattern, log, time_local)
						}
					case cat == "http_referer":
						if category == "Bad Referrer" {
							detect(options, category, pattern, log, time_local)
						}
					case cat == "request_uri":
						if category == "Directory Bruteforce" {
							detect(options, category, "^/"+pattern+"$", log, time_local)
						}
					}
				}
			}
		}
	}
}

func detect(o *common.Options, c string, p string, l string, t string) {
	e := o.Configs.Rules.Threat.Excludes
	m = matchers.IsMatch(p, l)
	if m {
		for x := 0; x < len(e); x++ {
			if e[x] == c {
				ex = true
			}
		}
		if !ex {
			gologger.Labelf("[%s] [%s] %s", t, c, l)
		}
	}
}
