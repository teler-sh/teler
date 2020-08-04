package configs

// Resources ...
type Resources struct {
	Threat []struct {
		Category string
		URL      string
		Content  string
		Exclude  bool
	}
}

var resource *Resources

// Init resources
func init() {
	resource = &Resources{
		Threat: []struct {
			Category string
			URL      string
			Content  string
			Exclude  bool
		}{
			{
				Category: "Common Web Attack",
				URL:      "https://raw.githubusercontent.com/PHPIDS/PHPIDS/master/lib/IDS/default_filter.json",
			},
			{
				Category: "Bad IP Address",
				URL:      "https://raw.githubusercontent.com/mitchellkrogza/nginx-ultimate-bad-bot-blocker/master/_generator_lists/bad-ip-addresses.list",
			},
			{
				Category: "Bad Referrer",
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
}

// Get default resources
func Get() *Resources {
	return resource
}
