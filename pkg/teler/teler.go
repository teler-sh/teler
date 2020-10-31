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
	"ktbs.dev/teler/pkg/metrics"
	"ktbs.dev/teler/resource"
)

// Analyze logs from threat resources
func Analyze(options *common.Options, logs *gonx.Entry) (bool, map[string]string) {
	var match bool
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

		log["category"] = cat
		metrics.GetThreasTotal.WithLabelValues(cat).Inc()

		if exc {
			continue
		}

		switch cat {
		case "Common Web Attack":
			req, err := url.Parse(log["request_uri"])
			if err != nil {
				break
			}
			query := req.Query()
			if len(query) > 0 {
				for p, q := range query {
					fil, _ := fastjson.Parse(con)
					dec, _ := url.QueryUnescape(strings.Join(q, ""))
					cwa := fil.GetArray("filters")

					for _, v := range cwa {
						log["category"] = cat + ": " + string(v.GetStringBytes("description"))
						log["element"] = "request_uri"
						quote := regexp.QuoteMeta(dec)

						if white := isWhitelist(options, p+"="+dec); white {
							continue
						}

						match = matchers.IsMatch(
							string(v.GetStringBytes("rule")),
							quote,
						)
						if match {
							metrics.GetCwa.WithLabelValues(string(v.GetStringBytes("description")),
								log["remote_addr"], log["request_uri"], log["status"]).Inc()

							break
						}
					}
				}
			}

		case "Bad Crawler":
			log["element"] = "http_user_agent"

			if white := isWhitelist(options, log["http_user_agent"]); white {
				break
			}

			for _, pat := range strings.Split(con, "\n") {
				if match = matchers.IsMatch(pat, log["http_user_agent"]); match {

					metrics.GetBadCrawler.WithLabelValues(log["remote_addr"],
						log["http_user_agent"],
						log["status"]).Inc()

					break
				}
			}

		case "Bad IP Address":
			log["element"] = "remote_addr"

			if white := isWhitelist(options, log["remote_addr"]); white {
				break
			}

			ip := "(?m)^" + log["remote_addr"]
			match = matchers.IsMatch(ip, con)
			metrics.GetBadIP.WithLabelValues(log["remote_addr"]).Inc()

		case "Bad Referrer":
			log["element"] = "http_referer"
			if white := isWhitelist(options, log["http_referer"]); white {
				break
			}
			if log["http_referer"] == "-" {
				break
			}

			req, _ := url.Parse(log["http_referer"])
			ref := "(?m)^" + req.Host

			match = matchers.IsMatch(ref, con)
			metrics.GetBadReferrer.WithLabelValues(log["http_referer"]).Inc()

		case "Directory Bruteforce":
			log["element"] = "request_uri"

			if white := isWhitelist(options, log["request_uri"]); white {
				break
			}

			req, err := url.Parse(log["request_uri"])
			if err != nil {
				break
			}

			if req.Path != "/" {
				match = matchers.IsMatch(trimFirst(req.Path), con)
			}

			if match {
				switch log["status"] {
				case "200", "204", "304":
					match = false
				}
				metrics.GetDirBruteforce.WithLabelValues(log["remote_addr"],
					log["request_uri"],
					log["status"]).Inc()
			}
		}

		if match {
			return match, log
		}
	}

	return match, log
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
