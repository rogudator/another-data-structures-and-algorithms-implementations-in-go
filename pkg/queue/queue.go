package queue

import "errors"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("Queue is empty")
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove, nil
}