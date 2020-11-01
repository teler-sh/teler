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
	Path  string
	file  string
	now   string
	cache Cache
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

// Check if resources is cached
func Check() bool {
	if err := configdir.MakePath(Path); err != nil {
		return false
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	fh, err := os.Open(file)
	if err != nil {
		return false
	}
	defer fh.Close()

	decoder := json.NewDecoder(fh)
	if err := decoder.Decode(&cache); err != nil {
		return false
	}

	updated, err := time.Parse(time.RFC3339, cache.UpdatedAt)
	if err != nil {
		return false
	}

	if time.Since(updated).Hours() < 24 {
		return true
	}

	return false
}

// Update latest resources being cached
func Update() {
	cache = Cache{now}

	fh, err := os.Create(file)
	if err != nil {
		errors.Exit(err.Error())
	}
	defer fh.Close()

	encoder := json.NewEncoder(fh)

	// nolint:errcheck
	encoder.Encode(&cache)
}

// Purge local/cached resources
func Purge() {
	err := os.RemoveAll(Path)
	if err != nil {
		errors.Exit(err.Error())
	}
}
