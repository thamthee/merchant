package adapter

// pointerStringToString using only transition graph model to mDB model.
func pointerStringToString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
