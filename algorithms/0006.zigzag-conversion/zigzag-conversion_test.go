package Problem0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
	two int
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
		{
			p: param{
				one: "PAYPALISHIRING",
				two: 3,
			},
			a: answer{
				one: "PAHNAPLSIIGYIR",
			},
		},
		{
			p: param{
				one: "PAYPALISHIRING",
				two: 4,
			},
			a: answer{
				one: "PINALSIGYAHRPI",
			},
		},
		{
			p: param{
				one: "A",
				two: 1,
			},
			a: answer{
				one: "A",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, convert(p.one, p.two), "Enter:%v", p)
	}
}