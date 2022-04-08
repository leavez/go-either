package result

import "github.com/leavez/go-either"

type Type[T any] either.Type[T, error]

func New[T any](v T) Type[T] {
	return Type[T](either.NewLeft[T, error](v))
}

func Error[T any](err error) Type[T] {
	return Type[T](either.NewRight[T, error](err))
}

func (s Type[T]) Unwrap() (T, error) {
	l, r, _ := either.Type[T, error](s).Unwrap()
	return l, r
}

func (s Type[T]) IsError() bool {
	_, _, isLeft := either.Type[T, error](s).Unwrap()
	return !isLeft
}

func Map[T, U any](r Type[T], f func(T) U) Type[U] {
	return Type[U](either.MapLeft(either.Type[T, error](r), f))
}
func MapError[T any](r Type[T], f func(err error) error) Type[T] {
	return Type[T](either.MapRight(either.Type[T, error](r), f))
}

func (s Type[T]) MapError(f func(err error) error) Type[T] {
	return MapError(s, f)
}
