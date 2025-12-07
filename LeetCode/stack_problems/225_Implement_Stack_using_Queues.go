package stack_problems

type MyStack struct {
	Idx   int
	Store []int
}

func NewMyStack() MyStack {
	return MyStack{Idx: -1}
}

func (s *MyStack) Push(x int) {
	s.Store = append(s.Store, x)
	s.Idx++
}

func (s *MyStack) Pop() int {
	elem := -1
	if s.Idx >= 0 {
		elem = s.Store[s.Idx]
		s.Store = s.Store[:s.Idx]
		s.Idx--
	}
	return elem
}

func (s *MyStack) Top() int {
	elem := -1
	if s.Idx >= 0 {
		elem = s.Store[s.Idx]
	}
	return elem
}

func (s *MyStack) Empty() bool {
	return !(s.Idx >= 0)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
