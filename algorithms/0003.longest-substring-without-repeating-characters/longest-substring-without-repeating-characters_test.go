package Problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
}

type answer struct {
	one int
}

type question struct {
	p param
	a answer
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: "abcabcbb",
			},
			a: answer{
				one: 3,
			},
		},
		{
			p: param{
				one: "bbbbbbbb",
			},
			a: answer{
				one: 1,
			},
		},
		{
			p: param{
				one: "pwwkew",
			},
			a: answer{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLongestSubstring(p.one), "Enter:%v", p)
	}
}