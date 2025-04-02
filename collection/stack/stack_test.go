package stack_test

import (
	"github.com/lewis-buji/go-utils/collection/stack"
	"testing"
)

func TestStack(t *testing.T) {
	stack1 := stack.New[int]()
	stack1.Push(1)
	stack1.Push(2)
	t.Log(stack1.Pop())
	t.Log(stack1.Pop())
	t.Log(stack1.Pop())
	t.Log(stack1.Pop())
	t.Log("===============================分割线===============================")

	type Person struct {
		Name string `json:"name"`
	}
	p1 := Person{Name: "p1"}
	p2 := Person{Name: "p2"}
	stack2 := stack.New[Person]()
	stack2.Push(p1)
	stack2.Push(p2)
	t.Log(stack2.Pop())
	t.Log(stack2.Pop())
	t.Log(stack2.Pop())
	t.Log("===============================分割线===============================")

	stack3 := stack.New[*Person]()
	stack3.Push(&p1)
	stack3.Push(&p2)
	t.Log(stack3.Pop())
	t.Log(stack3.Pop())
	t.Log(stack3.Pop())
}
