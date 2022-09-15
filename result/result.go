package result

import "errors"

type Result[T any] struct {
	Ok  T
	Err error
}

func Ok[T any](data T) Result[T] {
	return Result[T]{Ok: data, Err: nil}
}

func Err[T any](msg string) Result[T] {
	return Result[T]{Err: errors.New(msg)}
}
