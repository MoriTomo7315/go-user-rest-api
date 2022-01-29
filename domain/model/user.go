package model

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Prefecture string `json:"prefecture"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}
