package array

import "testing"

func TestOldReverse(t *testing.T) {
	s := []int{1, 2, 3, 4}
	n := OldReverse(s)
	t.Log(n)
}

func TestReverse(t *testing.T) {
	// ss := []int{1, 2, 3, 4}
	a := [...]int{0, 1, 2, 3, 4, 5} // 数组表示
	Reverse(a[:])
	t.Log(a)
}

func TestRunStack(t *testing.T) {
	RunStack()
	t.Log("ok")
}
