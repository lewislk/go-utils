package gptr_test

import (
	"github.com/lewis-buji/go-utils/gptr"
	"testing"
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
