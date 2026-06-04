import (
    "slices"
)

type MinStack struct {
    vals []int
}

func Constructor() MinStack {
    return MinStack{
        vals: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.vals = append(this.vals, val)
}

func (this *MinStack) Pop() {
    if len(this.vals) == 0 {
        return
    }
    this.vals = this.vals[:len(this.vals) - 1]
}

//This is one where I would change the signature to func(this *MinStack) Top() int, error {...}
func (this *MinStack) Top() int {
    return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
    return slices.Min(this.vals)
}
