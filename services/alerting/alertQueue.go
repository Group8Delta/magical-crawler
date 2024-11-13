package alerting

import (
	"container/heap"
)

type IAlertQueue interface {
	heap.Interface
}

type Alert struct {
	Title    string
	Content  string
	Priority int
	Index    int
}

type AlertQueue []*Alert

func (pq AlertQueue) Len() int { return len(pq) }

func (pq AlertQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq AlertQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *AlertQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Alert)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *AlertQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func NewAlertQueue() IAlertQueue {
	mq := &AlertQueue{}
	heap.Init(mq)
	return mq
}
