package messagingservice

import (
	"encoding/json"
	"log"
	"runtime"
	"sync"
	"test/mymom/infrastructure/crh"
	"test/mymom/services/messagingservice/event"
	"test/shared"
)

type MessagingService struct {
	Queues map[string]*MutexQueue
}

func NewMessagingService() MessagingService {
	qs := make(map[string]*MutexQueue)
	r := MessagingService{Queues: qs}
	return r
}

func (ms *MessagingService) Publish(qId string, e event.Event) bool {
	r := true

	ms.NotificationEngine(qId, e)

	return r // TODO
}

func (ms *MessagingService) Consume(qId string) {

	_, exists := ms.Queues[qId]
	if !exists {
		q := NewMutexQueue(100)
		ms.Queues[qId] = q
	}

	for {
		_crhX := crh.NewCRH(shared.LocalHost, shared.CallBackPort)
		msg, err := json.Marshal(ms.Queues[qId].Pop())
		if err != nil {
			log.Fatal("Messaging:: Callback:: encode error:", err)
		}
		_crhX.Send(msg)
	}
	return
}

type MutexQueue struct {
	mu       sync.Mutex
	capacity int
	queue    []event.Event
}

func NewMutexQueue(capacity int) *MutexQueue {
	return &MutexQueue{
		mu:       sync.Mutex{},
		capacity: capacity,
		queue:    []event.Event{},
	}
}

func (s *MutexQueue) Push(i event.Event) {

	s.mu.Lock()
	for len(s.queue) == s.capacity {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	s.queue = append(s.queue, i)
	s.mu.Unlock()
}

func (s *MutexQueue) Pop() event.Event {

	s.mu.Lock()
	for len(s.queue) == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	r := s.queue[0]
	s.queue = s.queue[1:]
	s.mu.Unlock()

	return r
}

func (s MutexQueue) Size() int {
	return len(s.queue)
}

func (ms *MessagingService) NotificationEngine(qId string, e event.Event) {

	_, exists := ms.Queues[qId]

	if exists {
		ms.Queues[qId].Push(e)
	} else {
		q := NewMutexQueue(shared.MaxQueueSize)
		ms.Queues[qId] = q
	}
}
