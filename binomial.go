package binomial_theorem

import (
	"fmt"
	"strconv"
	"strings"
)

type binomial struct {
	x expression
	y expression
	exp int
}

func (b *binomial) Expand() []expression {
	es := []expression{}
	su := PascalTriangle(b.exp)
	for i, v := range su {
		x := expression{b.x.literals, b.x.value, (len(su) - 1) - i}
		y := expression{b.y.literals, b.y.value, i}
		mult := Multiply(Multiply(&x, &y), &expression{[]literal{}, v, 1})
		es = append(es, *mult)
	}
	return es
}

func (b *binomial) String() string {
	var s strings.Builder

	fmt.Fprintf(&s, "(")

	fmt.Fprintf(&s, "%s", b.x.String())

	if b.y.value < 0 {
		fmt.Fprintf(&s, "-")
	} else {
		fmt.Fprintf(&s, "+")
	}

	fmt.Fprintf(&s, "%s", b.y.String())

	fmt.Fprintf(&s, ")")

	if b.exp != 1 {
		fmt.Fprintf(&s, "^%d", b.exp)
	}

	return s.String()
}

func ParseBinomial(str string) *binomial {
	rgx := `^\((?P<m1>-?\d*[a-z]*)+(?P<m2>[\+,-]\d*[a-z]*)+\)\^(?P<exp>\d+)$`
	result := RegexGroup(rgx, str)

	exp, err := strconv.Atoi(result["exp"])

	if err != nil {
		panic(err)
	}

	return &binomial{
		x: *ParseExpression(result["m1"]),
		y: *ParseExpression(result["m2"]),
		exp: exp,
	}
}
