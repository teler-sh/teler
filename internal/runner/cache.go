package runner

import (
	"github.com/projectdiscovery/gologger"
	"teler.app/pkg/cache"
	"teler.app/pkg/errors"
)

func rmCache() {
	cache.Purge()
	gologger.Info().Msg("All local cached resources have been removed.")
	errors.Abort(9)
}
