// Package choose 封装三元表达式
package choose

func If[T any](cond bool, onTrue, onFalse T) T {
	if cond {
		return onTrue
	}
	return onFalse
}

type lazy[T any] func() T

func IfLazy[T any](cond bool, onTrue, onFalse lazy[T]) T {
	if cond {
		return onTrue()
	}
	return onFalse()
}
