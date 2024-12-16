package ds

import "container/heap"

type Item[T any] struct {
	value    T
	priority int
	index    int
}

func (item Item[T]) Value() T {
	return item.value
}

func (item Item[T]) Priority() int {
	return item.priority
}

type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use lesser than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func CreateItem[T any](value T, priority int) *Item[T] {
	return &Item[T]{value: value, priority: priority}
}
