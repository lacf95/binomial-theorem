package binomial_theorem

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

type expression struct {
	literals []literal
	value int
	exp int
}

func Multiply(a *expression, b *expression) *expression {
	expr := expression{literals: []literal{}, value: 0, exp: 1}
	a = a.Expand()
	b = b.Expand()
	expr.value = a.value * b.value
	expr.literals = Expand(append(a.literals, b.literals...))
	return &expr
}

func (expr *expression) Expand() *expression {
	if expr.exp == 0 {
		return &expression{literals: []literal{}, value: 1, exp: 1}
	}

	if expr.value == 0 {
		return &expression{literals: []literal{}, value: 0, exp: 1}
	}

	n := &expression{literals: expr.literals, value: expr.value, exp: expr.exp}
	for i := 1; i < expr.exp; i ++ {
		n.value *= expr.value
		n.literals = Expand(append(n.literals, expr.literals...))
		n.exp -= 1
	}
	return n
}

func (expr *expression) String() string {
	var s strings.Builder
	absValue := int(math.Abs(float64(expr.value)))

	if expr.exp != 1 && len(expr.literals) > 1 {
		fmt.Fprintf(&s, "(")
	}

	if expr.value < 0 {
		fmt.Fprintf(&s, "-")
	}

	if absValue != 1 || len(expr.literals) == 0 {
		fmt.Fprintf(&s, "%d", absValue)
	}

	ls := GroupByExponent(expr.literals)
	for _, lg := range ls {
		for _, l := range lg {
			if l.exp != 0 {
				fmt.Fprintf(&s, "%s", l.name)
			}
		}
		if lg[0].exp > 1 {
			fmt.Fprintf(&s, "^%d", lg[0].exp)
		}
	}

	if expr.exp != 1 && len(expr.literals) > 1 {
		fmt.Fprintf(&s, ")")
	}

	if expr.exp != 1 {
		fmt.Fprintf(&s, "^%d", expr.exp)
	}

	return s.String()
}

func ParseExpression(str string) *expression {
	rgx := `^(?P<value>[\+,-]?\d*)(?P<literal>[a-z]*)$`
	result := RegexGroup(rgx, str)

	v, err := strconv.Atoi(result["value"])
	if err != nil {
		if result["value"] == "" {
			v = 1
		} else if result["value"] == "-" {
			v = -1
		} else {
			panic(err)
		}
	}

	exp, err := strconv.Atoi(result["exp"])
	if err != nil {
		exp = 1
	}

	ls := []literal{}
	for _, l := range strings.Split(result["literal"], "") {
		ls = append(ls, literal{l, 1})
	}

	return &expression{
		literals: ls,
		value: v,
		exp: exp,
	}
}
