package utils

func FormatTemperature(isIced bool, categoryID int) string {
	if categoryID == 1 || categoryID == 2 {
		if isIced {
			return "Ice"
		}
		return "Hot"
	}

	// category_id 3,4,5,6 and others
	return "-"
}
