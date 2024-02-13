package runner

import (
	"github.com/kitabisa/teler/pkg/cache"
	"github.com/kitabisa/teler/pkg/errors"
	"github.com/projectdiscovery/gologger"
)

func rmCache() {
	cache.Purge()
	gologger.Info().Msg("All local cached resources have been removed.")
	errors.Abort(9)
}
