package tools_test

import (
	"testing"

	"github.com/lewislk/go-utils/tools"
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
