package metrics

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"ktbs.dev/teler/common"
)

// PrometheusInsert logs into metrics
func PrometheusInsert(options *common.Options, data map[string]string) {
	var counter prometheus.Counter
	cfg := options.Configs

	//Check if teler have customs rules
	if cfg.Rules.Threat.Customs != nil {
		for _, custom := range cfg.Rules.Threat.Customs {
			if strings.HasPrefix(data["category"], custom.Name) {
				for _, rule := range custom.Rules {
					counter = getCustomsRule.WithLabelValues(
						data["category"],
						rule.Element,
						data[rule.Element],
					)
					counter.Inc()

					//if the rule use "or" operator the metrics will only get the fist match
					if strings.EqualFold(custom.Condition, "or") {
						break
					}
				}
			}
		}
	}

	switch {
	case strings.HasPrefix(data["category"], "Common Web Attack"):
		counter = getCWA.WithLabelValues(
			data["category"],
			data["remote_addr"],
			data["request_uri"],
			data["status"],
		)
	case strings.HasPrefix(data["category"], "CVE-"):
		counter = getCVE.WithLabelValues(
			data["category"],
			data["remote_addr"],
			data["request_uri"],
			data["status"],
		)
	case data["category"] == "Bad Crawler":
		counter = getBadCrawler.WithLabelValues(
			data["remote_addr"],
			data["http_user_agent"],
			data["status"],
		)
	case data["category"] == "Bad IP Address":
		counter = getBadIP.WithLabelValues(
			data["remote_addr"],
		)
	case data["category"] == "Bad Referrer":
		counter = getBadReferrer.WithLabelValues(
			data["http_referer"],
		)
	case data["category"] == "Directory Bruteforce":
		counter = getDirBruteforce.WithLabelValues(
			data["remote_addr"],
			data["request_uri"],
			data["status"],
		)
	default:
		return
	}

	counter.Inc()
	getThreatTotal.WithLabelValues(data["category"]).Inc()
}
