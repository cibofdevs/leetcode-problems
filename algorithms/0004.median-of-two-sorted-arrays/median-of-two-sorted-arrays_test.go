package Problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
	two []int
}

type answer struct {
	one float64
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
				one: []int{1, 3},
				two: []int{2},
			},
			a: answer{
				one: 2,
			},
		},
		{
			p: param{
				one: []int{1, 3},
				two: []int{2, 4},
			},
			a: answer{
				one: 2.5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findMedianSortedArrays(p.one, p.two), "Enter:%v", p)
	}

	ast.Panics(func() { findMedianSortedArrays([]int{}, []int{}) }, "Median of empty slices, but no panic")
}