package event

import (
	"encoding/json"
	"fmt"
	"html/template"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/goji/httpauth"
	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/pkg/errors"
	"github.com/projectdiscovery/gologger"
	"github.com/r3labs/sse/v2"
	"goji.io"
	"goji.io/pat"
)

// Run SSE
func Run(options *common.Options, version string) *server {
	srv := sse.New()
	srv.CreateStream("teler")

	s := &server{
		server:  srv,
		version: version,
		options: options,
	}
	a := httpauth.AuthOptions{Realm: "teler"}

	mux := goji.NewMux()

	user := s.options.Configs.Dashboard.Username
	pass := s.options.Configs.Dashboard.Password
	if (user != "") && (pass != "") {
		a.User = user
		a.Password = pass
		mux.Use(httpauth.BasicAuth(a))
	}
	mux.HandleFunc(pat.Get("/*"), s.static)

	h := s.options.Configs.Dashboard.Host
	if h == "" {
		h = "127.0.0.1"
	}

	p := s.options.Configs.Dashboard.Port
	if p == 0 {
		p = 8080
	}

	go func() {
		err := http.ListenAndServe(fmt.Sprint(h, ":", strconv.Itoa(p)), mux) // nosemgrep
		if err != nil {
			errors.Exit(err.Error())
		}
	}()

	gologger.Info().Msgf(fmt.Sprint("Listening dashboard on http://", h, ":", p))

	return s
}

// Push event for stream
func (s *server) Push(data map[string]string) {
	message, _ := json.Marshal(data)

	s.server.Publish("teler", &sse.Event{
		Data: message,
	})
}

func (s *server) static(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path

	switch p {
	case "/":
		p = "/index.html"
	case "/events":
		s.server.ServeHTTP(w, r)
		return
	}

	t, e := template.ParseFS(res, filepath.Join("www", p))
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(e.Error())) // nosemgrep
		return
	}

	x := filepath.Ext(p)
	if m := mime.TypeByExtension(x); m != "" {
		w.Header().Set("Content-Type", m)
	}
	w.WriteHeader(http.StatusOK)

	data := map[string]interface{}{
		"TELER_VERSION":            s.version,
		"TELER_DASHBOARD_HOST":     s.options.Configs.Dashboard.Host,
		"TELER_DASHBOARD_PORT":     s.options.Configs.Dashboard.Port,
		"TELER_DASHBOARD_ENDPOINT": s.options.Configs.Dashboard.Endpoint,
	}

	if e := t.Execute(w, data); e != nil {
		_, _ = w.Write([]byte(e.Error())) // nosemgrep
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
