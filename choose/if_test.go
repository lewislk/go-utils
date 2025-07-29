package choose_test

import (
	"testing"

	"github.com/lewislk/go-utils/choose"
)

func TestIf(t *testing.T) {
	if ans := choose.If(true, 1, 2); ans != 1 {
		t.Errorf("If(true, 1, 2) = %v; want 1", ans)
	}
	if ans := choose.IfLazy(false, func() string {
		return "1"
	}, func() string {
		return "2"
	}); ans != "2" {
		t.Errorf("choose.IfLazy... = %v; want 2", ans)
	}
}
