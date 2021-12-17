package metrics

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusInsert logs into metrics
func PrometheusInsert(data map[string]string) {
	var counter prometheus.Counter

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
