package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/kirsle/configdir"
	"ktbs.dev/teler/pkg/errors"
)

// Cache defines resources cache file informations
type Cache struct {
	UpdatedAt string `json:"updated_at"`
}

var (
	// Path define its local user-level cache path
	Path   string
	file   string
	now    string
	cache  Cache
	cached bool
)

func init() {
	Path = configdir.LocalCache("teler-resources")
	file = filepath.Join(Path, ".cached.json")

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		errors.Exit(err.Error())
	}

	now = time.Now().In(loc).Format(time.RFC3339)
}

// IsCached to check if resources is in local cache
func IsCached() (cached bool) {
	if err := configdir.MakePath(Path); err != nil {
		errors.Exit(err.Error())
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return
	}

	fh, err := os.Open(file)
	if err != nil {
		return
	}
	defer fh.Close()

	decoder := json.NewDecoder(fh)
	if err := decoder.Decode(&cache); err != nil {
		return
	}

	updated, err := time.Parse(time.RFC3339, cache.UpdatedAt)
	if err != nil {
		return
	}

	if time.Since(updated).Hours() < 24 {
		cached = true
	}

	return
}

// Update to updating local cache
func Update() {
	cache = Cache{now}

	fh, err := os.Create(file)
	if err != nil {
		errors.Exit(err.Error())
	}
	defer fh.Close()

	encoder := json.NewEncoder(fh)
	encoder.Encode(&cache)
}
