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
	"ktbs.dev/teler/pkg/matchers"
	"ktbs.dev/teler/resource"
)

// Analyze logs from threat resources
func Analyze(options *common.Options, logs *gonx.Entry) (bool, map[string]string) {
	var match bool

	out := make(map[string]string)
	log := make(map[string]string)
	rsc := resource.Get()

	fields := reflect.ValueOf(logs).Elem().FieldByName("fields")
	for _, field := range fields.MapKeys() {
		log[field.String()] = fields.MapIndex(field).String()
	}

	for i := 0; i < len(rsc.Threat); i++ {
		threat := reflect.ValueOf(&rsc.Threat[i]).Elem()
		cat := threat.FieldByName("Category").String()
		con := threat.FieldByName("Content").String()
		exc := threat.FieldByName("Exclude").Bool()

		out["date"] = log["time_local"]
		out["category"] = cat

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
						out["category"] += ": " + string(v.GetStringBytes("description"))
						out["element"] = log["request_uri"]
						quote := regexp.QuoteMeta(dec)

						if white := isWhitelist(options, quote); white {
							break
						}

						match = matchers.IsMatch(
							string(v.GetStringBytes("rule")),
							quote,
						)
						if match {
							break
						}
					}
				}
			}
		case "Bad Crawler":
			out["element"] = log["http_user_agent"]
			if white := isWhitelist(options, log["http_user_agent"]); white {
				break
			}

			for _, pat := range strings.Split(con, "\n") {
				if match = matchers.IsMatch(pat, log["http_user_agent"]); match {
					break
				}
			}
		case "Bad IP Address":
			out["element"] = log["remote_addr"]
			if white := isWhitelist(options, log["remote_addr"]); white {
				break
			}

			ip := "(?m)^" + log["remote_addr"]
			match = matchers.IsMatch(ip, con)
		case "Bad Referrer":
			out["element"] = log["http_referer"]
			if white := isWhitelist(options, log["http_referer"]); white {
				break
			}
			if log["http_referer"] == "-" {
				break
			}

			req, _ := url.Parse(log["http_referer"])
			ref := "(?m)^" + req.Path

			if req.Host != "" {
				ref += req.Host
			}

			match = matchers.IsMatch(ref, con)
		case "Directory Bruteforce":
			out["element"] = log["request_uri"]
			if white := isWhitelist(options, log["request_uri"]); white {
				break
			}

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

func isWhitelist(options *common.Options, find string) bool {
	whitelist := options.Configs.Rules.Threat.Whitelists
	for i := 0; i < len(whitelist); i++ {
		match := matchers.IsMatch(whitelist[i], find)
		if match {
			return true
		}
	}

	return false
}
