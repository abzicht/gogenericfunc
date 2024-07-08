package fun

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	type A int
	type E error

	var l Either[E, A]

	testErr := EitherError{"left error"}
	l = NewLeft[E, A](testErr)
	assert.Equal(t, testErr, l.GetLeft(), "Error in either is not correctly returned")
	assert.Panics(t, func() { _ = l.GetRight() }, "Expected left to panic at request for right")
	assert.True(t, cmp.Equal(l, NewLeft[E, A](EitherError{"left error"})))
	assert.False(t, cmp.Equal(l, NewRight[E, A](0)))
}

func TestRight(t *testing.T) {
	type A struct {
		A int
		B string
	}
	type E error

	var r Either[E, A]

	a := A{5, "test string"}

	r = NewRight[E, A](a)
	assert.Equal(t, a, r.GetRight(), "Value in either is not correctly returned")
	assert.Panics(t, func() { _ = r.GetLeft() }, "Expected right to panic at request for left")
	assert.True(t, cmp.Equal(r, NewRight[E, A](A{5, "test string"})))
	assert.False(t, cmp.Equal(r, NewLeft[E, A](EitherError{"an error"})))

	flattenedR := EitherFlatMap[E, A, A](r, func(a A) Either[E, A] {
		if a.A < 6 {
			return NewLeft[E, A](EitherError{"smaller than 6"})
		}
		return r
	})
	switch flattenedR.(type) {
	case Left[E, A]:
		break
	default:
		t.Error("Expected Left")
	}
	flattenedR = EitherFlatMap[E, A, A](r, func(a A) Either[E, A] {
		if a.A < 5 {
			return NewLeft[E, A](EitherError{"smaller than 5"})
		}
		return r
	})
	switch flattenedR.(type) {
	case Right[E, A]:
		break
	default:
		t.Error("Expected Right")
	}
}
