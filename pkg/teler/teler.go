package teler

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/satyrius/gonx"
	"github.com/valyala/fastjson"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/configs"
	"ktbs.dev/teler/pkg/matchers"
)

// Analyze logs from threat resources
func Analyze(options *common.Options, logs *gonx.Entry) (bool, map[string]string) {
	var match bool

	out := make(map[string]string)
	log := make(map[string]string)
	resource := configs.Get()

	fields := reflect.ValueOf(logs).Elem().FieldByName("fields")
	for _, field := range fields.MapKeys() {
		log[field.String()] = fields.MapIndex(field).String()
	}

	for i := 0; i < len(resource.Threat); i++ {
		threat := reflect.ValueOf(&resource.Threat[i]).Elem()
		cat := threat.FieldByName("Category").String()
		con := threat.FieldByName("Content").String()
		exc := threat.FieldByName("Exclude").Bool()

		out["date"] = log["time_local"]

		if exc {
			continue
		}

		switch cat {
		case "Common Web Attack":
			req, _ := url.Parse(log["request_uri"])
			query := req.Query()
			if len(query) > 0 {
				for _, q := range query {
					fil, _ := fastjson.Parse(con)
					dec, _ := url.QueryUnescape(strings.Join(q, ""))
					cwa := fil.GetArray("filters")

					for _, v := range cwa {
						match = matchers.IsMatch(
							string(v.GetStringBytes("rule")),
							regexp.QuoteMeta(dec),
						)
						out["category"] = cat + ": " + string(v.GetStringBytes("description"))
						out["element"] = log["request_uri"]

						if match {
							break
						}
					}
				}
			}
		case "Bad Crawler":
			out["category"] = cat
			out["element"] = log["http_user_agent"]

			for _, pat := range strings.Split(con, "\n") {
				match = matchers.IsMatch(pat, log["http_user_agent"])
				if match {
					break
				}
			}
		case "Bad IP Address":
			out["category"] = cat
			out["element"] = log["remote_addr"]

			ip := "(?m)^" + log["remote_addr"]
			match = matchers.IsMatch(ip, con)
		case "Bad Referrer":
			out["category"] = cat
			out["element"] = log["http_referer"]

			if log["http_referer"] == "-" {
				break
			}

			req, _ := url.Parse(log["http_referer"])
			ref := "(?m)^"
			ref += req.Path

			if req.Host != "" {
				ref += req.Host
			}

			match = matchers.IsMatch(ref, con)
		case "Directory Bruteforce":
			out["category"] = cat
			out["element"] = log["request_uri"]

			req, _ := url.Parse(log["request_uri"])

			if req.Path != "/" {
				match = matchers.IsMatch(trimFirst(req.Path), con)
			}
		}

		if match {
			return match, out
		}
	}
	return match, out
}

func trimFirst(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
