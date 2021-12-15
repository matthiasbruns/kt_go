package standard

func let[T any, R any](t *T, block func(t *T) R) R {
	return block(t)
}

func also[T  any](t *T, block func(t *T)) *T {
	block(t)
	return t
}

func takeIf[T any](t *T, predicate func(t *T) bool) *T {
	if predicate(t) {
		return t
	}
	return nil
}

func takeUnless[T  any](t *T, predicate func(t *T) bool) *T {
	if !predicate(t) {
		return t
	}
	return nil
}

func repeat(times int, action func(i int)) {
	for i := 0; i < times; i++ {
		action(i)
	}
}
