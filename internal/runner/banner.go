package runner

import (
	"github.com/projectdiscovery/gologger"
	"ktbs.dev/teler/common"
)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Printf("%s\n\n", common.Banner)
	gologger.Printf("\t%s\n\n", common.Email)
	if common.Development {
		gologger.Labelf("This tool is under development!")
		gologger.Labelf("Please submit a report if an error occurs.")
	}
}
