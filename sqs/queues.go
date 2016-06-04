package sqs

import (
	"fmt"
	"sync"
)

var (
	queuesMu sync.Mutex
	queues   []*Queue
)

func Queues() []*Queue {
	defer queuesMu.Unlock()
	queuesMu.Lock()
	return queues
}

func registerQueue(queue *Queue) {
	defer queuesMu.Unlock()
	queuesMu.Lock()

	for _, q := range queues {
		if q.name == queue.name {
			panic(fmt.Sprintf("%s is already registered", queue))
		}
	}
	queues = append(queues, queue)
}