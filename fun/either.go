package fun

import "github.com/google/go-cmp/cmp"

type Either[A, E any] interface {
	GetLeft() E
	GetRight() A
	FlatMap(func(A) Either[A, E]) Either[A, E]
	//Filter(func(A) bool) Either[A, E]
	Equal(Either[A, E]) bool
}

type EitherError struct {
	Msg string
}

func (ee EitherError) Error() string {
	return ee.Msg
}

type Left[A, E any] struct {
	e E
}

func NewLeft[A, E any](e E) Either[A, E] {
	return Left[A, E]{e}
}

func (l Left[A, E]) GetLeft() E                                  { return l.e }
func (l Left[A, E]) GetRight() A                                 { panic(EitherError{"Requested right value from a left object"}) }
func (l Left[A, E]) FlatMap(f func(A) Either[A, E]) Either[A, E] { return l }
func (l Left[A, E]) Filter(f func(A) bool) Either[A, E]          { return l }
func (l Left[A, E]) Equal(e Either[A, E]) bool {
	return cmp.Equal(l.GetLeft(), e.GetLeft()) && cmp.Equal(l.GetRight(), e.GetRight())
}

type Right[A, E any] struct {
	a A
}

func NewRight[A, E any](a A) Either[A, E] {
	return Right[A, E]{a}
}

func (r Right[A, E]) GetLeft() E  { panic(EitherError{"Requested left value from a right object"}) }
func (r Right[A, E]) GetRight() A { return r.a }
func (r Right[A, E]) FlatMap(f func(A) Either[A, E]) Either[A, E] {
	return f(r.GetRight())
}
func (r Right[A, E]) Equal(e Either[A, E]) bool {
	return cmp.Equal(r.GetLeft(), e.GetLeft()) && cmp.Equal(r.GetRight(), e.GetRight())
}
