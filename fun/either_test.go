package fun

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	type A int
	type E error

	var l Either[A, E]

	testErr := EitherError{"left error"}
	l = NewLeft[A, E](testErr)
	assert.Equal(t, testErr, l.GetLeft(), "Error in either is not correctly returned")
	assert.Panics(t, func() { _ = l.GetRight() }, "Expected left to panic at request for right")
	assert.True(t, cmp.Equal(l, NewLeft[A, E](EitherError{"left error"})))
	assert.False(t, cmp.Equal(l, NewRight[A, E](0)))
}

func TestRight(t *testing.T) {
	type A struct {
		A int
		B string
	}
	type E error

	var r Either[A, E]

	a := A{5, "test string"}

	r = NewRight[A, E](a)
	assert.Equal(t, a, r.GetRight(), "Value in either is not correctly returned")
	assert.Panics(t, func() { _ = r.GetLeft() }, "Expected right to panic at request for left")
	assert.True(t, cmp.Equal(r, NewRight[A, E](A{5, "test string"})))
	assert.False(t, cmp.Equal(r, NewLeft[A, E](EitherError{"an error"})))

	flattenedR := r.FlatMap(func(a A) Either[A, E] {
		if a.A < 6 {
			return NewLeft[A, E](EitherError{"smaller than 6"})
		}
		return r
	})
	switch flattenedR.(type) {
	case Left[A, E]:
		break
	default:
		t.Error("Expected Left")
	}
	flattenedR = r.FlatMap(func(a A) Either[A, E] {
		if a.A < 5 {
			return NewLeft[A, E](EitherError{"smaller than 5"})
		}
		return r
	})
	switch flattenedR.(type) {
	case Right[A, E]:
		break
	default:
		t.Error("Expected Right")
	}
}
