type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack {
		stack: []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		currentMin := this.GetMin()
		if currentMin > val {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, currentMin)
		}
	}
	
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
