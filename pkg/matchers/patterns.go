package matchers

const (
	PatternToken     = `^(xox[p|b|o|a]-\d{10,12}-\d{12}-\w+)|(\d{9}:[a-zA-Z0-9_-]{35})$`
	PatternChannel   = `^[A-Z0-9]{9,12}$`
	PatternHexcolor  = `^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`
	PatternLogformat = `\$(remote_(addr|user)|request_(method|uri|protocol)|http_(referer|user_agent))`
)
