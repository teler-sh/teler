package metrics

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// Send logs to metrics
func Send(log map[string]string) {
	var counter prometheus.Counter

	switch {
	case strings.HasPrefix(log["category"], "Common Web Attack"):
		counter = getCWA.WithLabelValues(
			log["category"],
			log["remote_addr"],
			log["request_uri"],
			log["status"],
		)
	case strings.HasPrefix(log["category"], "CVE-"):
		counter = getCVE.WithLabelValues(
			log["category"],
			log["remote_addr"],
			log["request_uri"],
			log["status"],
		)
	case log["category"] == "Bad Crawler":
		counter = getBadCrawler.WithLabelValues(
			log["remote_addr"],
			log["http_user_agent"],
			log["status"],
		)
	case log["category"] == "Bad IP Address":
		counter = getBadIP.WithLabelValues(
			log["remote_addr"],
		)
	case log["category"] == "Bad Referrer":
		counter = getBadReferrer.WithLabelValues(
			log["http_referer"],
		)
	case log["category"] == "Directory Bruteforce":
		counter = getDirBruteforce.WithLabelValues(
			log["remote_addr"],
			log["request_uri"],
			log["status"],
		)
	}

	counter.Inc()
	getThreatTotal.WithLabelValues(log["category"]).Inc()
}
