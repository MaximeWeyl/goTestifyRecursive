package formatting

import "strings"

//FormatFieldName Formats the field name. Use this to return nice errors in your custom behaviours
func FormatFieldName(parentString string, suffixes ...string) string {
	out := parentString
	suffixesString := strings.Join(suffixes, ".")

	if len(out) > 0 && len(suffixesString) > 0 {
		out += "."
	}
	out += suffixesString

	return out
}
