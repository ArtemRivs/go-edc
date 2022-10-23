package solution

import (
	"strings"
)

func Greetings(name string) string {
	return "Привет, " + strings.Title(strings.ToLower(strings.Trim(name, " "))) + "!"
}
