package stack_problems

type MyQueue struct {
	Queue []int
}

func NewMyQueue() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.Queue = append(q.Queue, x)
}

func (q *MyQueue) Pop() int {
	elem := -1
	if len(q.Queue) >= 0 {
		elem = q.Queue[0]
		q.Queue = q.Queue[1:]
	}
	return elem
}

func (q *MyQueue) Peek() int {
	elem := -1
	if len(q.Queue) >= 0 {
		elem = q.Queue[0]
	}
	return elem
}

func (this *MyQueue) Empty() bool {

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
