package solution

import (
	"strings"
)

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)
	return "Привет, " + name + "!"
}
