package errors

const (
	ErrCheckConfig    = "please check your config file"
	ErrConfigValidate = "Only validates :key; " + ErrCheckConfig
	ErrAlertProvider  = "Provider \":platform\" not available; " + ErrCheckConfig
	ErrNoInputLog     = "No input logs provided"
	ErrNoInputConfig  = "No config file specified"
	ErrNoFilePath     = "No file path specified; " + ErrCheckConfig
	ErrNoUserZinc     = "No username provided for Zinc log server; " + ErrCheckConfig
	ErrNoPassZinc     = "No password provided for Zinc log server; " + ErrCheckConfig
	ErrNoIndexZinc    = "No index provided for Zinc log server; " + ErrCheckConfig
	ErrAuthZinc       = "Invalid Zinc credentials"
	ErrHealthZinc     = "Zinc log server is not running"
	ErrInsertLogZinc  = "Failed to insert logs to Zinc"
)
