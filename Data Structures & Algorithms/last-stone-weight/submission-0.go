type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	*h = stones
	heap.Init(h)

	for h.Len() > 1 {
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)

		if first != second {
			heap.Push(h, first-second)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}