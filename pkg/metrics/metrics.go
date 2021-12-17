package metrics

import "github.com/prometheus/client_golang/prometheus"

// Defines its Prometheus metrics variables
var (
	getCWA = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_cwa",
			Help: "Get lists of Common Web Attack threats",
		},
		[]string{"description", "remote_addr", "request_uri", "status"},
	)

	getCVE = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_cve",
			Help: "Get lists of CVE threats",
		},
		[]string{"description", "remote_addr", "request_uri", "status"},
	)

	getBadCrawler = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_badcrawler",
			Help: "Get lists of Bad Crawler requests",
		},
		[]string{"remote_addr", "http_user_agent", "status"},
	)

	getDirBruteforce = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_dir_bruteforce",
			Help: "Get lists of Directories Bruteforced",
		},
		[]string{"remote_addr", "request_uri", "status"},
	)

	getBadIP = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_badip_count",
			Help: "Total number of Bad IP Addresses",
		},
		[]string{"remote_addr"},
	)

	getBadReferrer = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_bad_referrer",
			Help: "Get lists of Bad Referrer requests",
		},
		[]string{"http_referer"},
	)

	getThreatTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_threats_count_total",
			Help: "Total number of detected threats",
		},
		[]string{"case"},
	)
)

// Prometheus will register a metrics with the specified variables
func Prometheus() {
	prometheus.MustRegister(
		getBadCrawler, getDirBruteforce, getBadIP,
		getCWA, getCVE, getBadReferrer, getThreatTotal,
	)
}
