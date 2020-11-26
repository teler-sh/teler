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

		if exc {
			continue
		}

		switch cat {
		case "Common Web Attack":
			req, err := url.ParseRequestURI(log["request_uri"])
			if err != nil {
				break
			}

			query := req.Query()
			if len(query) > 0 {
				for p, q := range query {
					dec, err := url.QueryUnescape(strings.Join(q, ""))
					if err != nil {
						continue
					}

					if isWhitelist(options, p+"="+dec) {
						continue
					}

					cwa, _ := fastjson.Parse(con)
					for _, v := range cwa.GetArray("filters") {
						log["category"] = cat + ": " + string(v.GetStringBytes("description"))
						log["element"] = "request_uri"
						quote := regexp.QuoteMeta(dec)

						match = matchers.IsMatch(
							string(v.GetStringBytes("rule")),
							quote,
						)

						if match {
							metrics.GetCWA.WithLabelValues(
								log["category"],
								log["remote_addr"],
								log["request_uri"],
								log["status"],
							).Inc()

							break
						}
					}
				}
			}
		case "CVE":
			// var status bool

			req, err := url.ParseRequestURI(log["request_uri"])
			if err != nil {
				break
			}

			if isWhitelist(options, req.RequestURI()) {
				break
			}

			log["element"] = "request_uri"
			cves, _ := fastjson.Parse(con)
			for _, cve := range cves.GetArray("templates") {
				log["category"] = strings.ToTitle(string(cve.GetStringBytes("id")))

				for _, r := range cve.GetArray("requests") {
					method := string(r.GetStringBytes("method"))
					if method != log["request_method"] {
						continue
					}

					// for _, m := range r.GetArray("matchers") {
					// 	for _, s := range m.GetArray("status") {
					// 		if log["status"] == s.String() {
					// 			status = true
					// 		}
					// 	}
					// }

					// if !status {
					// 	continue
					// }

					for _, p := range r.GetArray("path") {
						diff, err := url.ParseRequestURI(
							strings.TrimPrefix(
								strings.Trim(p.String(), `"`),
								"{{BaseURL}}",
							),
						)
						if err != nil {
							continue
						}

						if len(diff.Path) <= 1 {
							continue
						}

						if req.Path != diff.Path {
							break
						}

						fq := 0
						for q := range req.Query() {
							if diff.Query().Get(q) != "" {
								fq++
							}
						}

						if len(diff.Query())-fq <= 3 {
							match = true
						}

						if match {
							metrics.GetCVE.WithLabelValues(
								log["category"],
								log["remote_addr"],
								log["request_uri"],
								log["status"],
							).Inc()

							break
						}
					}
				}

				if match {
					break
				}
			}
		case "Bad Crawler":
			log["element"] = "http_user_agent"

			if isWhitelist(options, log["http_user_agent"]) {
				break
			}

			for _, pat := range strings.Split(con, "\n") {
				if match = matchers.IsMatch(pat, log["http_user_agent"]); match {
					metrics.GetBadCrawler.WithLabelValues(
						log["remote_addr"],
						log["http_user_agent"],
						log["status"],
					).Inc()

					break
				}
			}
		case "Bad IP Address":
			log["element"] = "remote_addr"

			if isWhitelist(options, log["remote_addr"]) {
				break
			}

			ip := "(?m)^" + log["remote_addr"]
			match = matchers.IsMatch(ip, con)
			if match {
				metrics.GetBadIP.WithLabelValues(log["remote_addr"]).Inc()
			}
		case "Bad Referrer":
			log["element"] = "http_referer"
			if isWhitelist(options, log["http_referer"]) {
				break
			}

			if log["http_referer"] == "-" {
				break
			}

			req, err := url.Parse(log["http_referer"])
			if err != nil {
				break
			}
			ref := "(?m)^" + req.Host

			match = matchers.IsMatch(ref, con)
			if match {
				metrics.GetBadReferrer.WithLabelValues(log["http_referer"]).Inc()
			}
		case "Directory Bruteforce":
			log["element"] = "request_uri"

			if isWhitelist(options, log["request_uri"]) ||
				matchers.IsMatch("^20(0|4)$", log["status"]) ||
				matchers.IsMatch("^3[0-9]{2}$", log["status"]) {
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
				metrics.GetDirBruteforce.WithLabelValues(
					log["remote_addr"],
					log["request_uri"],
					log["status"],
				).Inc()
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
