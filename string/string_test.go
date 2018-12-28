package string

import "testing"

func TestReverseString(t *testing.T) {
	s := ReverseString("hello,wwwwa")
	t.Log(s)

	a := ReverseAllString("你好hello,wwwwa")
	t.Log(a)
}

func TestString2Int(t *testing.T) {
	i := String2Int("1234")
	t.Log(i)
}
