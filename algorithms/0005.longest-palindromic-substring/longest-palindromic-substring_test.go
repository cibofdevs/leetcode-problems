package Problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
}

type answer struct {
	one string
}

type question struct {
	p param
	a answer
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: param{
				one: "babad",
			},
			a: answer{
				one: "bab",
			},
		},
		question{
			p: param{
				one: "cbbd",
			},
			a: answer{
				one: "bb",
			},
		},
		question{
			p: param{
				one: "abbcccddcccbba",
			},
			a: answer{
				one: "abbcccddcccbba",
			},
		},
		question{
			p: param{
				one: "a",
			},
			a: answer{
				one: "a",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, longestPalindrome(p.one), "Enter:%v", p)
	}
}