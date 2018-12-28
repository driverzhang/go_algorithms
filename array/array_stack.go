/**************** push、pop 数组实现一个栈 *********************/
package array

import (
	"fmt"
	"sync"
)

type Stack struct {
	sync.RWMutex       // 考虑并发性
	Items        []int // 切片数组
	N            int   // 栈的长度
}

func NewStack(n int) *Stack {
	return &Stack{
		N:     n,
		Items: make([]int, 0, n),
	}
}

func (s *Stack) Push(i int, w *sync.WaitGroup) {
	s.Lock()
	defer s.Unlock()
	defer w.Done()
	if s.N == len(s.Items) { // 满栈情况
		return
	}

	s.Items = append(s.Items, i)
}

func (s *Stack) Pop(w *sync.WaitGroup) []int {
	s.Lock()
	defer s.Unlock()
	if len(s.Items) == 0 {
		return nil
	}

	s.Items = s.Items[:len(s.Items)-1]
	return s.Items
}

func RunStack() {
	var wg sync.WaitGroup
	stack := NewStack(10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go stack.Push(i, &wg)
	}

	wg.Wait()
	fmt.Println("old: ", stack.Items)

	nowStack := stack.Pop(&wg)
	fmt.Println("pop now: ", nowStack)
}
