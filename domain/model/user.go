package model

import "time"

type User struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Prefecture string    `json:"prefecture"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
