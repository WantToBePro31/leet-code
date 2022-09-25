type MyCircularQueue struct {
	Size int
	Count int
    Last int
	Elements []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Size: k,
		Count: 0,
        Last: 0,
		Elements: make([]int, k), 
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
    this.Elements[(this.Count+this.Last) % this.Size] = value
	this.Count++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
        this.Last = (this.Last + 1) % this.Size
		this.Count--
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if !this.IsEmpty() {
		return this.Elements[this.Last]
	}
	return -1
}

func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		return this.Elements[(this.Count+this.Last-1) % this.Size]
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.Count == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.Count == this.Size {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */