package phpgo

// If Golang版本的三元表达式，使用方法：If(a>b,a,b).(int)
func If(isTrue bool, a, b interface{}) interface{} {
	if isTrue {
		return a
	}
	return b
}
