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

	rules := make(map[string]bool)
	for _, custom := range cfg.Rules.Threat.Customs {
		rules[custom.Name] = true
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
	case rules[data["category"]]:
		counter = getCustomsRule.WithLabelValues(
			data["category"],
			data["element"],
			data[data["element"]],
		)
	default:
		return
	}

	counter.Inc()
	getThreatTotal.WithLabelValues(data["category"]).Inc()
}
