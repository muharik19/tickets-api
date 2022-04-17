package helper

func StringNullable(value interface{}) string {

	if value != nil {
		if len(value.(string)) <= 0 {
			return ""
		}
	} else {
		return ""
	}

	return value.(string)
}
