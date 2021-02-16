package common

var (
	Email       = "infosec@kitabisa.com"
	Development = true
	Version     = ""
	Banner      = `
	  __      __       
	 / /____ / /__ ____
	/ __/ -_) / -_) __/
	\__/\__/_/\__/_/   
	                ` + Version

	Usage = `  [buffers] | teler [options]
  teler -c [...] [options]`
	Example = `  teler -c [...] -i /var/log/nginx/access.log
  [kubectl logs|tail -f|cat] ... | teler -c [...] -x 25
  `
)
