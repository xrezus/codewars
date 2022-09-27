package convert_boolean_values

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}

	return "No"
}
