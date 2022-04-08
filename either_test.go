package either

import (
	"testing"
)

func assert(t *testing.T, condition bool, message string, args ...any) bool {
	c := !condition
	if c {
		t.Helper()
		t.Errorf(message, args...)
	}
	return c
}

func TestAll(t *testing.T) {
	e := NewLeft[string, int]("12")
	left, _, isLeft := e.Unwrap()
	if assert(t, isLeft, "should be true") {
		assert(t, left == "12", "")
	}

	e2 := MapLeft(e, func(left string) *string {
		return &left
	})
	left2, _, isLeft2 := e2.Unwrap()
	if assert(t, isLeft2, "should be true") {
		assert(t, *left2 == "12", "")
	}

	v := Map(e2, func(left *string) bool {
		return false
	}, func(right int) bool {
		assert(t, false, "shouldn't in this case")
		return true
	})
	assert(t, v == false, "")

	hit := false
	e.EitherDo(func(left string) {
		hit = true
	}, func(right int) {
		assert(t, false, "shouldn't in this case")
		hit = false
	})
	assert(t, hit, "")
}
