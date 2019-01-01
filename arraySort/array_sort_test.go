package arraySort

import "testing"

const n = 6

var a = []int{3, 5, 4, 1, 2, 6}

func TestBubbleSort(t *testing.T) {
	BubbleSort(a, n)
	t.Log(a)
}

func TestInsertionSort(t *testing.T) {
	a := []int{5, 4, 6, 1, 3, 2}

	InsertionSort(a, 6)
	t.Log(a)
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(a, n)
	t.Log(a)
}
