package subscription

func isValidFreqUnit(unit string) bool {
	switch unit {
	case "day", "week", "month", "year":
		return true
	default:
		return false
	}
}
