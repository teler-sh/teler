package runner

import "ktbs.dev/teler/versioninfo"

var (
	email       = "infosec@kitabisa.com"
	development = false
	banner      = `
	  __      __       
	 / /____ / /__ ____
	/ __/ -_) / -_) __/
	\__/\__/_/\__/_/   
	                ` + versioninfo.Version

	usage = `  [buffers] | teler [options]
  teler -c [...] [options]`
	example = `  teler -c [...] -i /var/log/nginx/access.log
  [kubectl logs|tail -f|cat] ... | teler -c [...] -x 25
  `
)
