package binomial_theorem

import(
	"fmt"
	"strings"
)

func BinomialTheorem(sexpr string) string {
	b := ParseBinomial(sexpr)

	return String(b.Expand())
}

func String(es []expression) string {
	var s strings.Builder

	for i, e := range es {
		if e.value >= 0 && i > 0 {
			fmt.Fprint(&s, "+")
		}
		fmt.Fprint(&s, e.String())
	}

	return s.String()
}
