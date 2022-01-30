package model

type ResponseModel struct {
	Status         int32 `json:"status"`
	Message       string `json:"message"`
	UserCount int64 `json:"user_count"`
	Users []*User `json:"users"`
}
