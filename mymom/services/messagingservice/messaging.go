package messagingservice

import (
	"encoding/json"
	"log"
	"test/mymom/infrastructure/crh"
	"test/mymom/services/messagingservice/event"
	"test/mymom/services/messagingservice/queue"
	"test/shared"
)

type MessagingService struct {
	Queues map[string]*queue.MutexQueue
}

func NewMessagingService() MessagingService {
	qs := make(map[string]*queue.MutexQueue)
	r := MessagingService{Queues: qs}
	return r
}

func (ms *MessagingService) Publish(qId string, e event.Event) bool {
	r := true

	ms.NotificationEngine(qId, e)

	return r // TODO
}

func (ms *MessagingService) Consume(qId string) {

	// Check if the queue exists
	_, exists := ms.Queues[qId]
	if !exists {
		q := queue.NewMutexQueue(shared.MaxQueueSize)
		ms.Queues[qId] = q
	}

	// Consume events from the queue
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

func (ms *MessagingService) NotificationEngine(qId string, e event.Event) {

	_, exists := ms.Queues[qId]

	if exists {
		ms.Queues[qId].Push(e)
	} else {
		q := queue.NewMutexQueue(shared.MaxQueueSize)
		ms.Queues[qId] = q
	}
}
