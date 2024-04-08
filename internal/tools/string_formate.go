package tools

func StringFormat(str string) string {
	short := str[1 : len(str)-1]
	return short
}
