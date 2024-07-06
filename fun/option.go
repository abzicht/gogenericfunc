package fun

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

/*
Options either hold "Some" value or "None". It is a fancy way to force you to check for NILs.
Given an option var "v" that may hold an int, find out, if it holds a value like so:

	switch v.(type) {
		case Some[int]:
			return v.GetValue()
		case None[int]:
			panic("No value in v :(")
	}

	Define new options that hold a value via NewSome(value) and placeholders via NewNone()
*/
type Option[T interface{}] interface {
	GetValue() T
	GetOrElse(T) T
	FlatMap(func(T) Option[T]) Option[T]
	Filter(func(T) bool) Option[T]
	Equal(Option[T]) bool
}

type OptionError struct {
	Msg string
}

func (oe OptionError) Error() string {
	return oe.Msg
}

type Some[T interface{}] struct {
	t T
}

func NewSome[T interface{}](t T) Some[T] {
	return Some[T]{t}
}

func NewNone[T interface{}]() None[T] {
	return None[T]{}
}

type None[T interface{}] struct {
}

func Try[T interface{}](f func() (T, error)) Option[T] {
	t, err := f()
	if err != nil {
		return None[T]{}
	}
	return Some[T]{t}
}

func (s Some[T]) GetValue() T {
	return s.t
}

func (n None[T]) GetValue() T {
	panic("This is seriously wrong and should never be reached.")
}

func (s Some[T]) GetOrElse(t T) T {
	return s.t
}

func (n None[T]) GetOrElse(t T) T {
	return t
}

func (s Some[T]) String() string {
	return fmt.Sprintf("[Some (Optional)] %v", s.t)
}

func (n None[T]) String() string {
	return "[None (Optional)]"
}

func (s Some[T]) Equal(o Option[T]) bool {
	switch o.(type) {
	case Some[T]:
		return cmp.Equal(s.GetValue(), o.GetValue())
	default:
		return false
	}
}

func (n None[T]) Equal(o Option[T]) bool {
	switch o.(type) {
	case None[T]:
		return true
	default:
		return false
	}
}

func OptionMap[T, U interface{}](opt Option[T], f func(t T) U) Option[U] {
	switch opt.(type) {
	case Some[T]:
		return Some[U]{f(opt.GetValue())}
	default:
		return None[U]{}
	}
}

func (s Some[T]) FlatMap(f func(t T) Option[T]) Option[T] {
	return f(s.t)
}

func (n None[T]) FlatMap(f func(t T) Option[T]) Option[T] {
	return n
}

func (s Some[T]) Filter(f func(t T) bool) Option[T] {
	if f(s.t) {
		return s
	}
	return None[T]{}
}

func (n None[T]) Filter(f func(t T) bool) Option[T] {
	return n
}
