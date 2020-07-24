package runner

import "github.com/projectdiscovery/gologger"

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Printf("%s\n", Banner)
	gologger.Printf("\t%s\n\n", Email)
	gologger.Labelf("This tool is under development.\n")
	gologger.Labelf("Please send a report if an error occurs.\n")
}
