package runner

import "github.com/projectdiscovery/gologger"

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Printf("%s\n\n", banner)
	gologger.Printf("\t%s\n\n", email)
	gologger.Labelf("This tool is under development!\n")
	gologger.Labelf("Please submit a report if an error occurs.\n")
}
