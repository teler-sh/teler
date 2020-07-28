package runner

const (
	email   = "secinfo@kitabisa.com"
	version = "0.0.3-dev2"
	banner  = `
	  __      __       
	 / /____ / /__ ____
	/ __/ -_) / -_) __/
	\__/\__/_/\__/_/   
	                v` + version

	usage = `  [buffers] | teler [options]
  teler -f [...] [options]`
	example = `  teler -f [...] -i /var/log/nginx/access.log
  [kubectl logs|tail -f|cat] ... | teler -f [...] -c 25`
)
