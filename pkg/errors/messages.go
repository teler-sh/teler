package errors

const (
	ErrCheckConfig          = "please check your config file"
	ErrConfigValidate       = "Only validates :key; " + ErrCheckConfig
	ErrNotificationProvider = "Provider \":platform\" not available; " + ErrCheckConfig
)
