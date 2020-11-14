package matchers

const (
	PatternToken     = `^(xox[p|b|o|a]-\d{10,12}-\d{12,13}-\w+)|(\d{9}:[a-zA-Z0-9_-]{35})|([MN][A-Za-z\d]{23}\.[\w-]{6}\.[\w-]{27})$`
	PatternChannel   = `^([A-Z0-9]{9,18})$`
	PatternColor     = `^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})|(\d{1,8})$`
	PatternLogformat = `\$(remote_(addr|user)|request_(method|uri|protocol)|http_(referer|user_agent))`
)
