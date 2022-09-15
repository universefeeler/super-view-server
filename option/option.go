package option

type Option[T any] struct {
	None bool
	Some T
}

func None[T any]() Option[T] {
	return Option[T]{None: true}
}

func Some[T any](data T) Option[T] {
	return Option[T]{None: false, Some: data}
}
