package result

import (
	"fmt"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	r := New("123")
	if w, err := r.Unwrap(); err != nil {
		assert(t, w == "123", "")
	}

	assert(t, r.IsError() == false, "")
	r2 := Map(r, func(t string) int {
		return 123
	})
	assert(t, r2.IsError() == false, "")
	if w, err := r2.Unwrap(); err != nil {
		assert(t, w == 123, "")
	}

	r3 := Error[int](fmt.Errorf("111"))
	if assert(t, r3.IsError(), "") {
		_, err := r.Unwrap()
		assert(t, err != nil, "")
	}
	r4 := MapError(r3, func(err error) error {
		return fmt.Errorf("failed :%w", err)
	})
	if assert(t, r4.IsError(), "") {
		_, err := r.Unwrap()
		assert(t, strings.Contains(err.Error(), "failed"), "")
	}

	assert(t, r == r.MapError(func(err error) error {
		return fmt.Errorf("failed :%w", err)
	}), "")
}

func assert(t *testing.T, condition bool, message string, args ...any) bool {
	c := !condition
	if c {
		t.Helper()
		t.Errorf(message, args...)
	}
	return c
}
