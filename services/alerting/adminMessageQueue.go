package alerting

import (
	"container/heap"
)

type IAdminMessageQueue interface {
	heap.Interface
}

type AdminMessage struct {
	Title    string
	Content  string
	Priority int
	Index    int
}

type AdminMessageQueue []*AdminMessage

func (pq AdminMessageQueue) Len() int { return len(pq) }

func (pq AdminMessageQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq AdminMessageQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *AdminMessageQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*AdminMessage)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *AdminMessageQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func NewAdminMessageQueue() IAdminMessageQueue {
	mq := &AdminMessageQueue{}
	heap.Init(mq)
	return mq
}
