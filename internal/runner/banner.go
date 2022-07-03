package runner

import (
	"github.com/projectdiscovery/gologger"
	"teler.app/common"
)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n\n", common.Banner)
	gologger.Print().Msgf("\t%s\n\n", common.Email)
	if common.Development {
		gologger.Warning().Msg("This tool is under development!")
		gologger.Warning().Msg("Please submit a report if an error occurs.")
	}
}
