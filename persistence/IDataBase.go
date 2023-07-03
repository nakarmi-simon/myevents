package persistence

import "github.com/myevents/models"

type DatabaseHandler interface {
	AddEvent(models.Event) ([]byte, error)
	FindEvent([]byte) (models.Event, error)
	FindEventByName(string) (models.Event, error)
	FindAllAvailableEvents() ([]models.Event, error)
}
