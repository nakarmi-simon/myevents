package contracts

import "time"

type EventCreatedEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID string    `json:"locationID"`
	Start      time.Time `json:"start_time"`
	End        time.Time `json:"end_time"`
}

func (e *EventCreatedEvent) EventName() string {
	return "event.Created"
}
