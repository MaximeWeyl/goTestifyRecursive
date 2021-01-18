package formatting

import "strings"

func FormatFieldName(parentString string, suffixes ...string) string {
	out := parentString
	suffixesString := strings.Join(suffixes, ".")

	if len(out) > 0 && len(suffixesString) > 0 {
		out += "."
	}
	out += suffixesString

	return out
}
