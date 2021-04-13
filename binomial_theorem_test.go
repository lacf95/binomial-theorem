package binomial_theorem

import "testing"

func TestBinomialTheorem(t *testing.T) {
	type test struct {
		binomial string
		want string
	}

	tests := []test{
		{"(x+1)^0", "1"},
		{"(x+1)^1", "x+1"},
		{"(x+1)^2", "x^2+2x+1"},
		{"(x-1)^0", "1"},
		{"(x-1)^1", "x-1"},
		{"(x-1)^2", "x^2-2x+1"},
		{"(5m+3)^4", "625m^4+1500m^3+1350m^2+540m+81"},
		{"(2x-3)^3", "8x^3-36x^2+54x-27"},
		{"(7x-7)^0", "1"},
		{"(-5m+3)^4", "625m^4-1500m^3+1350m^2-540m+81"},
		{"(-2k-3)^3", "-8k^3-36k^2-54k-27"},
		{"(-7x-7)^0", "1"},
	}

    for _, tc := range tests {
        got := BinomialTheorem(tc.binomial)
        if tc.want != got {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}
