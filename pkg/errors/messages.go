package errors

const (
	ErrCheckConfig    = "please check your config file"
	ErrConfigValidate = "Only validates :key; " + ErrCheckConfig
	ErrAlertProvider  = "Provider \":platform\" not available; " + ErrCheckConfig
	ErrNoInputLog     = "No input logs provided"
	ErrNoInputConfig  = "No config file specified"
)
