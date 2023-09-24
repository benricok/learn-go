package main

import (
	"time"

	"github.com/google/uuid"
)

type CustomTime struct {
	time.Time
}

type Item struct {
	ID        uuid.UUID  `json:"id"`
	Owner     string     `json:"owner"`
	Name      string     `json:"name"`
	Done      bool       `json:"done"`
	Created   CustomTime `json:"datecreated,omitempty"`
	Completed CustomTime `json:"datecompleted,omitempty"`
}

type User struct {
	Username string     `json:"username"`
	Surname  string     `json:"surname"`
	Name     string     `json:"name"`
	Created  CustomTime `json:"datecreated"`
}