package Problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one *ListNode
	two *ListNode
}

type answer struct {
	one *ListNode
}

type question struct {
	p param
	a answer
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: makeListNode([]int{2,4,3}),
				two: makeListNode([]int{5,6,4}),
			},
			a: answer{
				one: makeListNode([]int{7,0,8}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{9,8,7,6,5}),
				two: makeListNode([]int{1,1,2,3,4}),
			},
			a: answer{
				one: makeListNode([]int{0,0,0,0,0,1}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{0}),
				two: makeListNode([]int{5,6,4}),
			},
			a: answer{
				one: makeListNode([]int{5,6,4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addTwoNumbers(p.one, p.two), "Enter:%v", p)
	}
}