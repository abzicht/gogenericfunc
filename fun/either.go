package fun

/*
Either holds either a value of type A or of type E.
*/
type Either[E, A any] interface {
	GetLeft() E
	GetRight() A
	//FlatMap(func(A) Either[E, A]) Either[E, A]
}

type EitherError struct {
	Msg string
}

func (ee EitherError) Error() string {
	return ee.Msg
}

type Left[E, A any] struct {
	E E
}

func NewLeft[E, A any](e E) Either[E, A] {
	return Left[E, A]{e}
}

func (l Left[E, A]) GetLeft() E  { return l.E }
func (l Left[E, A]) GetRight() A { panic(EitherError{"Requested right value from a left object"}) }

type Right[E, A any] struct {
	A A
}

func NewRight[E, A any](a A) Either[E, A] {
	return Right[E, A]{a}
}

func (r Right[E, A]) GetLeft() E  { panic(EitherError{"Requested left value from a right object"}) }
func (r Right[E, A]) GetRight() A { return r.A }

func EitherFlatMap[E, A, B any](either Either[E, A], mapFunc func(A) Either[E, B]) Either[E, B] {
	switch either.(type) {
	case Right[E, A]:
		return mapFunc(either.GetRight())
	case Left[E, A]:
		return NewLeft[E, B](either.GetLeft())
	}
	panic(EitherError{"Provided either is neither left nor right"})
}

func EitherMap[E, A, B any](either Either[E, A], mapFunc func(A) B) Either[E, B] {
	switch either.(type) {
	case Right[E, A]:
		return NewRight[E, B](mapFunc(either.GetRight()))
	case Left[E, A]:
		return NewLeft[E, B](either.GetLeft())
	}
	panic(EitherError{"Provided either is neither left nor right"})
}

func EitherTry[A any](tryFunc func() (A, error)) Either[error, A] {
	resultA, resultE := tryFunc()
	if resultE != nil {
		return NewLeft[error, A](resultE)
	}
	return NewRight[error, A](resultA)
}
