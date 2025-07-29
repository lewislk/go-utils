package gptr_test

import (
	"testing"

	"github.com/lewislk/go-utils/gptr"
)

func TestOf(t *testing.T) {
	gptr.Of(1)
	type People struct {
		Name string
		Age  int
	}
	people := People{Name: "lewis", Age: 18}
	peoplePtr := gptr.Of(people)
	t.Log(peoplePtr == &people) // 因为golang是值传递的，所以这两个指针是不同的
}
