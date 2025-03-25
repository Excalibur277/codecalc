package listenerstack

import "fmt"

type data interface{}

type ListenerStack []data

func (s *ListenerStack) Pop(amount int) (items ListenerQueue) {
	if len(*s) < amount {
		panic(fmt.Sprintf("listener expected %d values: had %d", amount, len(*s)))
	}
	var deq []data
	deq, *s = (*s)[len(*s)-amount:], (*s)[0:len(*s)-amount]
	items = make(ListenerQueue, amount)
	copy(items, deq)
	return items
}

func (s *ListenerStack) Push(d data) {
	*s = append(*s, d)
}

type ListenerQueue []data

func (s *ListenerQueue) Dequeue() (item data) {
	if len(*s) == 0 {
		panic("listener expected value: wasnt there")
	}
	item, *s = (*s)[0], (*s)[1:]
	return item
}
