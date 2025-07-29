package strs_test

import (
	"testing"

	"github.com/lewislk/go-utils/gptr"
	"github.com/lewislk/go-utils/strs"
)

func TestIsPtrEmpty(t *testing.T) {
	if ans := strs.IsPtrEmpty(gptr.Of("")); ans != true {
		t.Errorf("strs.IsPtrEmpty... = %v; want true", ans)
	}
	if ans := strs.IsPtrEmpty(nil); ans != true {
		t.Errorf("strs.IsPtrEmpty... = %v; want true", ans)
	}
	if ans := strs.IsPtrNumeric(gptr.Of("123")); ans != true {
		t.Errorf("strs.IsPtrNumeric... = %v; want true", ans)
	}
	if ans := strs.IsPtrBlank(gptr.Of("\r\n")); ans != true {
		t.Errorf("strs.IsPtrBlank... = %v; want true", ans)
	}
}
