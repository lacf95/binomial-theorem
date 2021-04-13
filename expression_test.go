package binomial_theorem

import "testing"

func TestExpression(t *testing.T) {
	type test struct {
        expr expression
        want string
    }

    tests := []test{
		{expression{literals: []literal{{"a", 1}}, value: -1, exp: 2}, "-a^2"},
		{expression{literals: []literal{{"b", 1}}, value: 1, exp: 2}, "b^2"},
		{expression{literals: []literal{{"x", 1}}, value: 2, exp: 1}, "2x"},
		{expression{literals: []literal{{"z", 1}}, value: 3, exp: 4}, "3z^4"},
		{expression{literals: []literal{}, value: 4, exp: 3}, "4^3"},
		{expression{literals: []literal{}, value: -5, exp: 0}, "-5^0"},
		{expression{literals: []literal{{"m", 1}}, value: 8, exp: 0}, "8m^0"},
		{expression{literals: []literal{{"m", 1}}, value: -1, exp: 1}, "-m"},
		{expression{literals: []literal{}, value: -1, exp: 1}, "-1"},
		{expression{literals: []literal{}, value: 7, exp: 1}, "7"},
		{expression{literals: []literal{{"h", 1}}, value: 0, exp: 2}, "0h^2"},
		{expression{literals: []literal{{"h", 1}}, value: 0, exp: 0}, "0h^0"},
		{expression{literals: []literal{{"x", 1}, {"y", 1}}, value: 3, exp: 1}, "3xy"},
		{expression{literals: []literal{{"x", 2}, {"y", 2}}, value: 3, exp: 1}, "3xy^2"},
		{expression{literals: []literal{{"x", 1}, {"y", 2}}, value: 3, exp: 1}, "3y^2x"},
		{expression{literals: []literal{{"x", 1}, {"y", 2}}, value: 3, exp: 2}, "(3y^2x)^2"},
	}

    for _, tc := range tests {
        got := tc.expr.String()
        if tc.want != got {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}

func TestMultiplyExpression(t *testing.T) {
	type test struct {
        a expression
        b expression
        want string
    }

    tests := []test{
		{expression{literals: []literal{}, value: 2, exp: 1}, expression{literals: []literal{}, value: 1, exp: 1}, "2"},
		{expression{literals: []literal{{"x", 1}}, value: 2, exp: 1}, expression{literals: []literal{{"y", 1}}, value: 1, exp: 1}, "2xy"},
		{expression{literals: []literal{{"x", 1}}, value: 2, exp: 1}, expression{literals: []literal{{"x", 1}, {"y", 1}}, value: 1, exp: 1}, "2x^2y"},
		{expression{literals: []literal{{"x", 2}}, value: 2, exp: 1}, expression{literals: []literal{{"x", 1}}, value: 2, exp: 1}, "4x^3"},
    }

    for _, tc := range tests {
        got := Multiply(&tc.a, &tc.b).String()
        if tc.want != got {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}
func TestExpandExpression(t *testing.T) {
	type test struct {
        expr expression
        want string
    }

    tests := []test{
		{expression{literals: []literal{{"a", 1}}, value: -1, exp: 2}, "a^2"},
		{expression{literals: []literal{{"b", 1}}, value: 1, exp: 2}, "b^2"},
		{expression{literals: []literal{{"x", 1}}, value: 2, exp: 1}, "2x"},
		{expression{literals: []literal{{"z", 1}}, value: 3, exp: 4}, "81z^4"},
		{expression{literals: []literal{}, value: 4, exp: 3}, "64"},
		{expression{literals: []literal{}, value: -5, exp: 0}, "1"},
		{expression{literals: []literal{{"m", 1}}, value: 8, exp: 0}, "1"},
		{expression{literals: []literal{{"m", 1}}, value: -1, exp: 1}, "-m"},
		{expression{literals: []literal{}, value: -1, exp: 1}, "-1"},
		{expression{literals: []literal{}, value: 7, exp: 1}, "7"},
		{expression{literals: []literal{{"h", 1}}, value: 0, exp: 2}, "0"},
		{expression{literals: []literal{{"h", 1}}, value: 0, exp: 0}, "1"},
		{expression{literals: []literal{{"x", 1}, {"y", 1}}, value: 3, exp: 1}, "3xy"},
		{expression{literals: []literal{{"x", 2}, {"y", 2}}, value: 3, exp: 1}, "3xy^2"},
		{expression{literals: []literal{{"x", 1}, {"y", 2}}, value: 3, exp: 1}, "3y^2x"},
		{expression{literals: []literal{{"x", 1}, {"y", 2}}, value: 3, exp: 2}, "9y^4x^2"},
    }

    for _, tc := range tests {
        got := tc.expr.Expand().String()
        if tc.want != got {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}
