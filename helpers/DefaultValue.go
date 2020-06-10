package helpers

func DefaultValue(a string, defaultVal string) string {
	if a != "" {
		return a
	} else {
		return defaultVal
	}
}
