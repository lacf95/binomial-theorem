package binomial_theorem

import "sort"

type literal struct {
	name string
	exp int
}

func Expand(ls []literal) []literal {
	nls := []literal{}
	for _, lg := range GroupByLiteral(ls) {
		newL := literal{lg[0].name, 0}
		for _, l := range lg {
			newL.exp += l.exp
		}
		nls = append(nls, newL)
	}
	return nls
}

func GroupByExponent(ls []literal) [][]literal {
	res := make(map[int][]literal)
	for _, l := range ls {
		res[l.exp] = append(res[l.exp], l)
	}

	as_arr := [][]literal{}
	for _, lg := range res {
		as_arr = append(as_arr, lg)
	}

	sort.Sort(byExponent(as_arr))
	return as_arr
}

func GroupByLiteral(ls []literal) [][]literal {
	res := make(map[string][]literal)
	for _, l := range ls {
		res[l.name] = append(res[l.name], l)
	}

	as_arr := [][]literal{}
	for _, lg := range res {
		as_arr = append(as_arr, lg)
	}

	sort.Sort(byExponent(as_arr))
	return as_arr
}

type byExponent [][]literal

func (ls byExponent) Len() int {
	return len(ls)
}

func (ls byExponent) Swap(a, b int) {
	ls[a], ls[b] = ls[b], ls[a]
}

func (ls byExponent) Less(a, b int) bool {
	return ls[a][0].exp > ls[b][0].exp
}
