package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Name string

type Event struct {
	ID bson.ObjectId `bson:"_id"`
	Name
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

type Location struct {
	Name
	ID        bson.ObjectId
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Halls
}

type Halls struct {
	Name     `json:"name"`
	Location string `json:"Location,omitempty"`
	Capacity string `json:"capacity"`
}
