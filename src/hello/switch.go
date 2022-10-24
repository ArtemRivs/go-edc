package solution

import ("strings")

func ModifySpaces(s, mode string) string {
	switch {

		case mode == "dash":
			s = strings.ReplaceAll(s, " ", "-")
		case mode == "underscore":
			s = strings.ReplaceAll(s, " ", "_")
		default:
			s = strings.ReplaceAll(s, " ", "*")

	}

	return s
}
