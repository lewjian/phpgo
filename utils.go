package phpgo

// If Golang版本的三元表达式
func If[T any](isTrue bool, a, b T) T {
	if isTrue {
		return a
	}
	return b
}
