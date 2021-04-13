package binomial_theorem

import "regexp"

func RegexGroup(expr string, str string) map[string]string {
	result := make(map[string]string)
	rgx := regexp.MustCompile(expr)

	match := rgx.FindStringSubmatch(str)

	for i, name := range rgx.SubexpNames() {
		if i != 0 && name != "" {
			if len(match) > i {
				result[name] = match[i]
			} else {
				result[name] = ""
			}
		}
	}
	return result
}
