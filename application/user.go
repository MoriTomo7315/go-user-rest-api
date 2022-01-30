package application

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	myError "github.com/MoriTomo7315/go-user-rest-api/domain/error"
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
)

// インターフェース
type UserApplication interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request, userId string)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request, userId string)
	DeleteUser(w http.ResponseWriter, r *http.Request, userId string)
}

type userApplication struct {
	firestoreRepository repository.FirestoreRepository
	responseRepository     repository.ResponseRepository
	errorRepository     repository.ErrorRepository
}

// Userデータに関するUseCaseを生成
func NewUserApplication(fr repository.FirestoreRepository, rr repository.ResponseRepository, er repository.ErrorRepository) UserApplication {
	return &userApplication{
		firestoreRepository: fr,
		responseRepository: rr,
		errorRepository:     er,
	}
}

// ユーザ一覧取得
func (ua userApplication) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ua.firestoreRepository.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	resModel := ua.responseRepository.GetResponse(http.StatusOK, "ユーザ情報取得に成功しました。", int64(len(users)), users)
	res, _ := json.Marshal(resModel)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ID指定でユーザ取得
func (ua userApplication) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {

	user, err := ua.firestoreRepository.GetUserById(userId)
	if err != nil {
		log.Printf("ERROR userが見つかりませんでした。 userId=%s", userId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		errorModel := ua.errorRepository.GetErrorResponse(myError.NOT_FOUND_USER_ERR_MSG)
		res, _ := json.Marshal(errorModel)
		w.Write(res)
	}
	users := []*model.User{user}
	resModel := ua.responseRepository.GetResponse(http.StatusOK, "ユーザ情報取得に成功しました。", int64(len(users)), users)
	res, _ := json.Marshal(resModel)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ユーザ作成
func (ua userApplication) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var user *model.User
	json.Unmarshal(body, &user)
	err := ua.firestoreRepository.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	resModel := ua.responseRepository.GetResponse(http.StatusCreated, "作成に成功しました。", 0, nil)
	res, _ := json.Marshal(resModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ユーザ更新
func (ua userApplication) UpdateUser(w http.ResponseWriter, r *http.Request, userId string) {

	// user存在 チェック
	_, err := ua.firestoreRepository.GetUserById(userId)
	if err != nil {
		log.Printf("ERROR userが見つかりませんでした。 userId=%s", userId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		errorModel := ua.errorRepository.GetErrorResponse(myError.NOT_FOUND_USER_ERR_MSG)
		res, _ := json.Marshal(errorModel)
		w.Write(res)
	}

	body, _ := ioutil.ReadAll(r.Body)

	var newUser *model.User
	json.Unmarshal(body, &newUser)
	newUser.Id = userId
	err = ua.firestoreRepository.UpdateUser(newUser)
	if err != nil {
		log.Fatal(err)
	}

	resModel := ua.responseRepository.GetResponse(http.StatusNoContent, "更新に成功しました。", 0, nil)
	res, _ := json.Marshal(resModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ユーザ削除
func (ua userApplication) DeleteUser(w http.ResponseWriter, r *http.Request, userId string) {
	// user存在 チェック
	_, err := ua.firestoreRepository.GetUserById(userId)
	if err != nil {
		log.Printf("ERROR userが見つかりませんでした。 userId=%s", userId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		errorModel := ua.errorRepository.GetErrorResponse(myError.NOT_FOUND_USER_ERR_MSG)
		res, _ := json.Marshal(errorModel)
		w.Write(res)
	}

	err = ua.firestoreRepository.DeleteUser(userId)
	if err != nil {
		log.Fatal(err)
	}
	resModel := ua.responseRepository.GetResponse(http.StatusNoContent, "削除に成功しました。", 0, nil)
	res, _ := json.Marshal(resModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
