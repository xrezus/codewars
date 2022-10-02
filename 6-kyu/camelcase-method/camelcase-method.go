package camelcase_method

import . "strings"

func CamelCase(s string) string {
	return ReplaceAll(Title(s), " ", "")
}
