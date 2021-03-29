package main

import (
	"log"
	"math/rand"
	"time"
)

// type used to enumerate events
type EventType string

const (
	event1 EventType = "event1"
	event2 EventType = "event2"
)

type Event struct {
	Type EventType
	Data interface{}
}

func eventSender(c chan Event) {

	for {
		// Send a rendom event to the channel
		rand.Seed(time.Now().Unix())
		events := []EventType{
			event1,
			event2,
		}
		n := rand.Int() % len(events)

		// send event to channel
		c <- Event{
			Type: events[n],
			Data: "test",
		}

		// wait a bit
		time.Sleep(2 * time.Second)
	}

}

// Create a struct to hold config
// And simplify dependency injections
type EventHandler struct {
	// Event channel
	Events chan Event
	// hold the registedred event functionss
	EventMap map[EventType][]func(Event)
}

func NewEventHandler() *EventHandler {
	eventMap := make(map[EventType][]func(Event))
	events := make(chan Event)

	return &EventHandler{
		Events:   events,
		EventMap: eventMap,
	}
}

// Register the handler function to handle an event type
func (h *EventHandler) Handle(e EventType, f func(Event)) {
	h.EventMap[e] = append(h.EventMap[e], f)
}

func (h *EventHandler) EventDispatcher() {
	for evt := range h.Events {
		log.Printf("event recieved: %v", evt)
		if handlers, ok := h.EventMap[evt.Type]; ok {
			// If we registered an event
			for _, f := range handlers {
				// exacute function as goroutine
				go f(evt)
			}
		}
	}
}

func main() {

	eventHandler := NewEventHandler()

	eventHandler.Handle(event1, func(e Event) {
		log.Printf("event Handled: %v", e)
	})

	go eventSender(eventHandler.Events)

	eventHandler.EventDispatcher()

}
