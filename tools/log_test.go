package tools_test

import (
	"github.com/lewis-buji/go-utils/tools"
	"testing"
)

func TestGetLogStr(t *testing.T) {
	t.Log(tools.GetLogStr(1))
	type People struct {
		Name string
		Age  int
	}
	people := People{Name: "lewis", Age: 18}
	t.Log(tools.GetLogStr(people))
}
