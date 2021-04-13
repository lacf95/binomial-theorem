package binomial_theorem

import "testing"

func TestExpandBinomial(t *testing.T) {
	type test struct {
        b binomial
        want []expression
    }
    tests := []test{
		{binomial{ // (x+2)^2
			expression{literals: []literal{{"x", 1}}, value: 1, exp: 1},
			expression{literals: []literal{}, value: 2, exp: 1},
			2,
		}, []expression{
			{literals: []literal{{"x", 2}}, value: 1, exp: 1},
			{literals: []literal{{"x", 1}}, value: 4, exp: 1},
			{literals: []literal{}, value: 4, exp: 1},
		}},
		{binomial{ // (x+2)^2
			expression{literals: []literal{{"x", 1}}, value: 1, exp: 1},
			expression{literals: []literal{}, value: 2, exp: 1},
			3,
		}, []expression{
			{literals: []literal{{"x", 3}}, value: 1, exp: 1},
			{literals: []literal{{"x", 2}}, value: 6, exp: 1},
			{literals: []literal{{"x", 1}}, value: 12, exp: 1},
			{literals: []literal{}, value: 8, exp: 1},
		}},
		{binomial{ // (x^2+2y^3)^2
			expression{literals: []literal{{"x", 2}}, value: 1, exp: 1},
			expression{literals: []literal{{"y", 3}}, value: 2, exp: 1},
			2,
		}, []expression{
			{literals: []literal{{"x", 4}}, value: 1, exp: 1},
			{literals: []literal{{"x", 2}, {"y", 3}}, value: 4, exp: 1},
			{literals: []literal{{"y", 6}}, value: 4, exp: 1},
		}},
		{binomial{ // (-x^2+2y^3)^2
			expression{literals: []literal{{"x", 2}}, value: -1, exp: 1},
			expression{literals: []literal{{"y", 3}}, value: 2, exp: 1},
			2,
		}, []expression{
			{literals: []literal{{"x", 4}}, value: 1, exp: 1},
			{literals: []literal{{"x", 2}, {"y", 3}}, value: -4, exp: 1},
			{literals: []literal{{"y", 6}}, value: 4, exp: 1},
		}},
    }

    for _, tc := range tests {
        got := tc.b.Expand()
        for i, expr := range got {
			if expr.String() != tc.want[i].String() {
				t.Errorf("expected: %v, got: %v", tc.want[i].String(), expr.String())
			}
        }
    }
}
