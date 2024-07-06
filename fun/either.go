package fun

/*
Either holds either a value of type A or of type E.
*/
type Either[A, E any] interface {
	GetLeft() E
	GetRight() A
	FlatMap(func(A) Either[A, E]) Either[A, E]
}

type EitherError struct {
	Msg string
}

func (ee EitherError) Error() string {
	return ee.Msg
}

type Left[A, E any] struct {
	E E
}

func NewLeft[A, E any](e E) Either[A, E] {
	return Left[A, E]{e}
}

func (l Left[A, E]) GetLeft() E                                  { return l.E }
func (l Left[A, E]) GetRight() A                                 { panic(EitherError{"Requested right value from a left object"}) }
func (l Left[A, E]) FlatMap(f func(A) Either[A, E]) Either[A, E] { return l }

type Right[A, E any] struct {
	A A
}

func NewRight[A, E any](a A) Either[A, E] {
	return Right[A, E]{a}
}

func (r Right[A, E]) GetLeft() E  { panic(EitherError{"Requested left value from a right object"}) }
func (r Right[A, E]) GetRight() A { return r.A }
func (r Right[A, E]) FlatMap(f func(A) Either[A, E]) Either[A, E] {
	return f(r.GetRight())
}
