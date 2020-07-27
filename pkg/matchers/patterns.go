package matchers

const (
	PatternToken     = `^(xox[p|b|o|a]-\d{10,12}-\d{12}-\w+)|(\d{9}:[a-zA-Z0-9_-]{35})$`
	PatternChannel   = `^[A-Z0-9]{9,12}$`
	PatternHexcolor  = `^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`
	PatternParseMode = `^(Markdown(?:V2)?|HTML)$`
)
