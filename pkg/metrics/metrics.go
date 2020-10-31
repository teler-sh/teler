package metrics

import "github.com/prometheus/client_golang/prometheus"

//metric variables
var (
	GetCwa = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_cwa",
			Help: "teler_cwa show all cwa threats",
		},
		[]string{"description", "remote_addr", "request_uri", "status"},
	)

	GetBadCrawler = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_badcrawler",
			Help: "teler_badcrawler get a bad crawler",
		},
		[]string{"remote_addr", "http_user_agent", "status"},
	)

	GetDirBruteforce = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_dir_bruteforce",
			Help: "teler_dir_bruteforce list of bruteforced dir",
		},
		[]string{"remote_addr", "request_uri", "status"},
	)

	GetBadIP = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_badip_count",
			Help: "get_badip remote address of threats",
		},
		[]string{"remote_addr"},
	)

	GetBadReferrer = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_bad_referrer",
			Help: "get a http with bad referrer",
		},
		[]string{"http_referer"},
	)

	GetThreasTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_threats_count_total",
			Help: "get threats count",
		},
		[]string{"case"},
	)
)

// MetricInit will register a metric with the specified variables
func MetricInit() {
	prometheus.MustRegister(GetBadCrawler,
		GetDirBruteforce, GetBadIP,
		GetCwa, GetBadReferrer, GetThreasTotal)
}
