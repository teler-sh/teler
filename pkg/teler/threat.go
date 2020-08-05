package teler

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/configs"
	"github.com/kitabisa/teler/pkg/matchers"
	"github.com/projectdiscovery/gologger"
	"github.com/satyrius/gonx"
	"github.com/valyala/fastjson"
)

// Analyze logs from threat resources
func Analyze(options *common.Options, logs *gonx.Entry) {
	var match bool
	log := make(map[string]string)
	resource := configs.Get()

	fields := reflect.ValueOf(logs).Elem().FieldByName("fields")
	for _, field := range fields.MapKeys() {
		log[field.String()] = fields.MapIndex(field).String()
	}

	go func() {
		for i := 0; i < len(resource.Threat); i++ {
			threat := reflect.ValueOf(&resource.Threat[i]).Elem()
			cat := threat.FieldByName("Category").String()
			con := threat.FieldByName("Content").String()
			exc := threat.FieldByName("Exclude").Bool()

			if exc {
				continue
			}

			switch cat {
			case "Common Web Attack":
				req, _ := url.Parse(log["request_uri"])
				query := req.Query()
				if len(query) > 0 {
					go func() {
						for _, q := range query {
							fil, _ := fastjson.Parse(con)
							dec, _ := url.QueryUnescape(strings.Join(q, ""))
							cwa := fil.GetArray("filters")
							for _, v := range cwa {
								match = matchers.IsMatch(string(v.GetStringBytes("rule")), regexp.QuoteMeta(dec))
								detect(match, cat+": "+string(v.GetStringBytes("description")), log["request_uri"], log["time_local"])
							}
						}
					}()
				}
			case "Bad Crawler":
				go func() {
					for _, pat := range strings.Split(con, "\n") {
						match = matchers.IsMatch(pat, log["http_user_agent"])
						if match {
							break
						}
					}
				}()
				detect(match, cat, log["http_user_agent"], log["time_local"])
			case "Bad IP Address":
				ip := "(?m)^" + log["remote_addr"]
				match = matchers.IsMatch(ip, con)
				detect(match, cat, log["remote_addr"], log["time_local"])
			case "Bad Referrer":
				ref := "(?m)^"
				if log["http_referer"] == "-" {
					continue
				}
				req, _ := url.Parse(log["http_referer"])
				ref += req.Path
				if req.Host != "" {
					ref += req.Host
				}
				match = matchers.IsMatch(ref, con)
				detect(match, cat, log["http_referer"], log["time_local"])
			case "Directory Bruteforce":
				req, _ := url.Parse(log["request_uri"])
				if matchers.IsMatch("^(2|3)[0-9]{2}", log["status"]) {
					continue
				}

				if req.Path != "/" {
					match = matchers.IsMatch(trimFirst(req.Path), con)
					detect(match, cat, log["request_uri"], log["time_local"])
				}
			}
		}
	}()
}

func trimFirst(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func detect(m bool, c string, l string, t string) {
	if m {
		gologger.Labelf("[%s] [%s] %s", t, c, l)
	}
}
