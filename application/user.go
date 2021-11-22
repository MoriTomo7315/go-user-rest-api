package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
)

// インターフェース
type UserApplication interface {
	// GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	// CreateUser(w http.ResponseWriter, r *http.Request)
	// UpdateUser(w http.ResponseWriter, r *http.Request)
	// DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userApplication struct {
	firestoreRepository repository.FirestoreRepository
	errorRepository     repository.ErrorRepository
}

// Userデータに関するUseCaseを生成
func NewUserApplication(fr repository.FirestoreRepository, er repository.ErrorRepository) UserApplication {
	return &userApplication{
		firestoreRepository: fr,
		errorRepository:     er,
	}
}

// ユーザ一覧取得
// func (ua userApplication) GetUsers(w http.ResponseWriter, r *http.Request) {

// }

// ID指定でユーザ取得
func (ua userApplication) GetUserById(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	user, err := ua.firestoreRepository.GetUserById(id)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ユーザ作成
// func (ua userApplication) CreateUser(w http.ResponseWriter, r *http.Request) {

// }

// // ユーザ更新
// func (ua userApplication) UpdateUser(w http.ResponseWriter, r *http.Request) {

// }

// // ユーザ削除
// func (ua userApplication) DeleteUser(w http.ResponseWriter, r *http.Request) {

// }