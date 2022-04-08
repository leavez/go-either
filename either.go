package either

import "github.com/leavez/go-optional"

type Type[T, U any] struct {
	left  optional.Type[T]
	right optional.Type[U]
}

func NewLeft[T, U any](v T) Type[T, U] {
	return Type[T, U]{
		left:  optional.New[T](v),
		right: optional.Nil[U](),
	}
}
func NewRight[T, U any](v U) Type[T, U] {
	return Type[T, U]{
		left:  optional.Nil[T](),
		right: optional.New[U](v),
	}
}

func (s Type[T, U]) Unwrap() (left T, right U, isLeft bool) {
	if l, ok := s.left.Value(); ok {
		left = l
		isLeft = true
		return
	}
	if r, ok := s.right.Value(); ok {
		right = r
		isLeft = false
		return
	}
	return
}

func (s Type[Left, Right]) EitherDo(leftCase func(Left), rightCase func(Right)) {
	if l, ok := s.left.Value(); ok {
		leftCase(l)
		return
	}
	if r, ok := s.right.Value(); ok {
		rightCase(r)
		return
	}
	panic("impossible")
}

func Map[Left, Right, R any](s Type[Left, Right], leftCase func(Left) R, rightCase func(Right) R) R {
	if l, ok := s.left.Value(); ok {
		return leftCase(l)
	}
	if r, ok := s.right.Value(); ok {
		return rightCase(r)
	}
	panic("impossible")
}

func MapLeft[Left, Right, LeftNew any](e Type[Left, Right], f func(Left) LeftNew) Type[LeftNew, Right] {
	if w, ok := e.left.Value(); ok {
		return NewLeft[LeftNew, Right](f(w))
	}
	return NewRight[LeftNew, Right](e.right.ForceValue())
}

func MapRight[Left, Right, RightNew any](e Type[Left, Right], f func(Right) RightNew) Type[Left, RightNew] {
	if w, ok := e.right.Value(); ok {
		return NewRight[Left, RightNew](f(w))
	}
	return NewLeft[Left, RightNew](e.left.ForceValue())
}
