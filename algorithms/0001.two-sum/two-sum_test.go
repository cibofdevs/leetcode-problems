package Problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
	two int
}

type answer struct {
	one []int
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
				one: []int{2,7,11,15},
				two: 9,
			},
			a: answer{
				one: []int{0,1},
			},
		},
		{
			p: param{
				one: []int{3,2,4},
				two: 8,
			},
			a: answer{
				one: nil,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoSum(p.one, p.two), "Enter:%v", p)
	}
}