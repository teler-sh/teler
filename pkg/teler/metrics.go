package teler

import "github.com/prometheus/client_golang/prometheus"

var (
	getcwa = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_get_cwa",
			Help: "get_cwa request.",
		},
		[]string{"description", "http_user_agent", "remote_addr", "request_uri", "status"},
	)

	getbadcrawler = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_get_badcrawler",
			Help: "get_badcrawler request.",
		},
		[]string{"remote_addr", "http_user_agent", "status"},
	)

	getdirbruteforce = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_get_dir_bruteforce",
			Help: "get_dir_bruteforce request.",
		},
		[]string{"remote_addr", "http_user_agent", "request_uri", "status"},
	)

	getbadip = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_get_badip",
			Help: "get_badip request.",
		},
		[]string{"remote_addr"},
	)

	getbadreferrer = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teler_get_referrer",
			Help: "get_referrer request.",
		},
		[]string{"http_referer"},
	)
)

//add var to register
func MetricInit() {
	prometheus.MustRegister(getbadcrawler,
		getdirbruteforce, getbadip,
		getcwa, getbadreferrer)
}
