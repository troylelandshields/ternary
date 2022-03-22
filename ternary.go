package ternary

type TernaryBuild[T any] struct {
	condition bool
	tResult   T
}

// Give starts the ternary chain, but a value will not be returned until `Else` is called.
// For example: Give(value1).If(condition).Else(value2)
func Give[T any](tResult T) TernaryBuild[T] {
	return TernaryBuild[T]{
		tResult: tResult,
	}
}

func (t TernaryBuild[T]) If(condition bool) TernaryBuild[T] {
	t.condition = condition
	return t
}

func (t TernaryBuild[T]) Else(fResult T) T {
	if t.condition {
		return t.tResult
	}
	return fResult
}

// Convenience Functions

// GiveOnSuccess can be used with any function that returns a value and a boolean "ok".
// If the ok indicates success, the result of the function will be returned, otherwise the Else value will be returned.
func GiveOnOK[T any](tResult T, ok bool) TernaryBuild[T] {
	return TernaryBuild[T]{
		tResult:   tResult,
		condition: ok,
	}
}

// GiveOnSuccess can be used with any function that returns a value and an error.
// If the error is nil, the result of the function will be returned, otherwise the Else value will be returned.
func GiveOnSuccess[T any](tResult T, err error) TernaryBuild[T] {
	return TernaryBuild[T]{
		tResult:   tResult,
		condition: err == nil,
	}
}
