package configs

// Resources ...
type Resources struct {
	Threat []struct {
		Category string
		URL      string
		Content  string
	}
}

// Init resources
func Init() *Resources {
	resource := &Resources{
		Threat: []struct {
			Category string
			URL      string
			Content  string
		}{
			{
				Category: "Common Web Attack",
				URL:      "https://raw.githubusercontent.com/enygma/expose/master/src/Expose/filter_rules.json",
			},
			{
				Category: "Bad IP Address",
				URL:      "https://raw.githubusercontent.com/mitchellkrogza/nginx-ultimate-bad-bot-blocker/master/_generator_lists/bad-ip-addresses.list",
			},
			{
				Category: "Bad Refferer",
				URL:      "https://raw.githubusercontent.com/mitchellkrogza/nginx-ultimate-bad-bot-blocker/master/_generator_lists/bad-referrers.list",
			},
			{
				Category: "Bad Crawler",
				URL:      "https://raw.githubusercontent.com/JayBizzle/Crawler-Detect/master/raw/Crawlers.txt",
			},
			{
				Category: "Directory Bruteforce",
				URL:      "https://raw.githubusercontent.com/maurosoria/dirsearch/master/db/dicc.txt",
			},
		},
	}

	return resource
}
