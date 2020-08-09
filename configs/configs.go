package configs

// Resources define external threats resource
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
	db := "https://github.com/kitabisa/teler-resources/raw/master/db"
	resource = &Resources{
		Threat: []struct {
			Category string
			URL      string
			Content  string
			Exclude  bool
		}{
			{
				Category: "Common Web Attack",
				URL:      db + "/common-web-attacks.json",
			},
			{
				Category: "Bad IP Address",
				URL:      db + "/bad-ip-addresses.txt",
			},
			{
				Category: "Bad Referrer",
				URL:      db + "/bad-referrers.txt",
			},
			{
				Category: "Bad Crawler",
				URL:      db + "/bad-crawlers.txt",
			},
			{
				Category: "Directory Bruteforce",
				URL:      db + "/directory-bruteforces.txt",
			},
		},
	}
}

// Get default resources
func Get() *Resources {
	return resource
}
