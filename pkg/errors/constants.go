package errors

const (
	ErrAlertProvider  = "Provider \":platform\" not available; " + ErrCheckConfig
	ErrAuthZinc       = "Invalid Zinc credentials"
	ErrBlankField     = ":field can't be blank; " + ErrCheckConfig
	ErrCheckConfig    = "please check your config file"
	ErrConfigValidate = "Only validates :key; " + ErrCheckConfig
	ErrDupeCategory   = "Duplicated name for ':category' threat category; " + ErrCheckConfig
	ErrHealthZinc     = "Zinc log server is not running"
	ErrInsertLogZinc  = "Failed to insert logs to Zinc"
	ErrNoElement      = "Can't find ':element' on log format for ':category' threat category; " + ErrCheckConfig
	ErrNoFilePath     = "No file path specified; " + ErrCheckConfig
	ErrNoIndexZinc    = "No index provided for Zinc log server; " + ErrCheckConfig
	ErrNoInputConfig  = "No config file specified"
	ErrNoInputLog     = "No input logs provided"
	ErrNoPassZinc     = "No password provided for Zinc log server; " + ErrCheckConfig
	ErrNoThreatRules  = "No rules for ':category' threat category; " + ErrCheckConfig
	ErrNoUserZinc     = "No username provided for Zinc log server; " + ErrCheckConfig
	ErrParseConfig    = "Can't parse config file: "
)
