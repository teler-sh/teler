package teler

import "github.com/prometheus/client_golang/prometheus"

var (
	getcwa = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_cwa",
			Help: "teler_cwa show all cwa threats",
		},
		[]string{"description", "remote_addr", "request_uri", "status"},
	)

	getbadcrawler = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_badcrawler",
			Help: "teler_badcrawler get a bad crawler",
		},
		[]string{"remote_addr", "http_user_agent", "status"},
	)

	getdirbruteforce = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_dir_bruteforce",
			Help: "teler_dir_bruteforce list of bruteforced dir",
		},
		[]string{"remote_addr", "request_uri", "status"},
	)

	getbadip = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_badip_count",
			Help: "get_badip remote address of threats",
		},
		[]string{"remote_addr"},
	)

	getbadreferrer = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_bad_referrer",
			Help: "get a http with bad referrer",
		},
		[]string{"http_referer"},
	)

	getthreatstotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_threats_count_total",
			Help: "get threats count",
		},
		[]string{"case"},
	)
)

//add var to register
func MetricInit() {
	prometheus.MustRegister(getbadcrawler,
		getdirbruteforce, getbadip,
		getcwa, getbadreferrer, getthreatstotal)
}
