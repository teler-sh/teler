package resource

// Resources define external threats resource
type Resources struct {
	Threat []struct {
		Category string
		Filename string
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
			Filename string
			Content  string
			Exclude  bool
		}{
			{
				Category: "Common Web Attack",
				Filename: "common-web-attacks.json",
			},
			{
				Category: "CVE",
				Filename: "cves.json",
			},
			{
				Category: "Bad IP Address",
				Filename: "bad-ip-addresses.txt",
			},
			{
				Category: "Bad Referrer",
				Filename: "bad-referrers.txt",
			},
			{
				Category: "Bad Crawler",
				Filename: "bad-crawlers.txt",
			},
			{
				Category: "Directory Bruteforce",
				Filename: "directory-bruteforces.txt",
			},
		},
	}
}

// Get default resources
func Get() *Resources {
	return resource
}
