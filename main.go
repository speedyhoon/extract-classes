package extractClassNames

import (
	"regexp"
)

//Extract class names & ids from a string
func Extract(input string) (list []string) {
	if len(input) == 0 {
		return
	}

	//Replace unwanted whitespace with a single space
	input = regexp.MustCompile(`[\0'"\\\n\r\v\t\f]`).ReplaceAllString(input, " ")

	chopStart := regexp.MustCompile(`(?mis)^[^\.#]*`)
	chopEnd := regexp.MustCompile("(?mis)[ ~!@$%^&*()+=,/';:\"?><[\\]{}|`].*")

	for _, rule := range splitAll(input) {
		rule = chopStart.ReplaceAllLiteralString(rule, "")
		rule = chopEnd.ReplaceAllLiteralString(rule, "")
		if rule != "" && rule != "#" && rule != "." {
			list = append(list, rule)
		}
	}
	return
}

//ExtractByte extracts class (or id) name from a byte slice
func ExtractByte(input []byte) []string {
	return Extract(string(input))
}

//splitAll is similar to regexp.Split & returns a slice of the substrings between & prefixed with "#" or "." characters
func splitAll(s string) []string {
	re := regexp.MustCompile("([.#])")

	matches := re.FindAllStringIndex(s, -1)
	strings := make([]string, 0, len(matches))

	var beg, end int

	for _, match := range matches {
		end = match[0]
		if match[1] != 0 {
			strings = append(strings, s[beg:end])
		}
		beg = match[1]
		if beg > 0 {
			//Go back one character to include "#" or "."
			beg--
		}
	}

	if end != len(s) {
		strings = append(strings, s[beg:])
	}

	return strings
}
