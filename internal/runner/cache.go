package runner

import (
	"github.com/projectdiscovery/gologger"
	"ktbs.dev/teler/pkg/cache"
	"ktbs.dev/teler/pkg/errors"
)

func rmCache() {
	cache.Purge()
	gologger.Infof("All local cached resources have been removed.")
	errors.Abort(9)
}
