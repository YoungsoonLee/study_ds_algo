package stackByQueues

type Queue struct {
	Data []int
}

func (queue *Queue) EnQueue(data int) {
	queue.Data = append(queue.Data, data)
}

func (queue *Queue) Front() int {
	if len(queue.Data) > 0 {
		return queue.Data[0]
	}
	return 0
}

func (queue *Queue) DeQueue() int {
	if len(queue.Data) > 0 {
		data := queue.Data[0]
		queue.Data = queue.Data[1:]
		return data
	}
	return 0
}

func (queue *Queue) Size() int {
	return len(queue.Data)
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.Data) == 0
}

type Stack struct {
	Q1 Queue
	Q2 Queue
}

func NewStack() Stack {
	return Stack{}
}

func (stack *Stack) Push(data int) {
	if stack.Q1.IsEmpty() {
		stack.Q2.EnQueue(data)
	} else {
		stack.Q1.EnQueue(data)
	}
}

func (stack *Stack) Pop() int {
	if stack.Q2.IsEmpty() {
		i, n := 1, stack.Q1.Size()
		for i < n && !stack.Q1.IsEmpty() {
			stack.Q2.EnQueue(stack.Q1.DeQueue())
			i++
		}
		return stack.Q1.DeQueue()
	}
	i, n := 1, stack.Q2.Size()
	for i < n && !stack.Q2.IsEmpty() {
		stack.Q1.EnQueue(stack.Q2.DeQueue())
		i++
	}
	return stack.Q2.DeQueue()
}

func (stack *Stack) Peek() int {
	if stack.Q2.IsEmpty() {
		i, n := 1, stack.Q1.Size()
		for i < n && !stack.Q1.IsEmpty() {
			stack.Q2.EnQueue(stack.Q1.DeQueue())
			i++
		}
		temp := stack.Q1.DeQueue()
		stack.Q2.EnQueue(temp)
		return temp
	}

	i, n := 1, stack.Q2.Size()
	for i < n && !stack.Q2.IsEmpty() {
		stack.Q1.EnQueue(stack.Q2.DeQueue())
		i++
	}
	temp := stack.Q2.DeQueue()
	stack.Q1.EnQueue(temp)
	return temp
}
