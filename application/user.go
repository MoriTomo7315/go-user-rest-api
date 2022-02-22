package application

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MoriTomo7315/go-user-rest-api/application/util"
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
	logger "github.com/MoriTomo7315/go-user-rest-api/gcplogger"
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
}

// Userデータに関するUseCaseを生成
func NewUserApplication(fr repository.FirestoreRepository) UserApplication {
	return &userApplication{
		firestoreRepository: fr,
	}
}

// ユーザ一覧取得
func (ua userApplication) GetUsers(w http.ResponseWriter, r *http.Request) {
	traceId := logger.GetTraceId(r)

	// firestore repositoryでのログでも使用するためcontextにtraceIdをセット
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	log.Printf(logger.InfoLogEntry("[GetUsers] Application logic start", traceId))
	users, err := ua.firestoreRepository.GetUsers(ctx)
	if err != nil {
		util.CreateErrorResponse(w, ctx, err, "")
		return
	}
	resModel := util.GetResponse(http.StatusOK, "ユーザ情報取得に成功しました。", int64(len(users)), users)
	res, _ := json.Marshal(resModel)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ID指定でユーザ取得
func (ua userApplication) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	traceId := logger.GetTraceId(r)
	// firestore repositoryでのログでも使用するためcontextにtraceIdをセット
	ctx := context.WithValue(context.Background(), "traceId", traceId)
	log.Printf(logger.InfoLogEntry("[GetUserById] Application logic start", traceId))
	user, err := ua.firestoreRepository.GetUserById(ctx, userId)
	if err != nil {
		util.CreateErrorResponse(w, ctx, err, userId)
		return
	}
	users := []*model.User{user}
	resModel := util.GetResponse(http.StatusOK, "ユーザ情報取得に成功しました。", int64(len(users)), users)
	res, _ := json.Marshal(resModel)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ユーザ作成
func (ua userApplication) CreateUser(w http.ResponseWriter, r *http.Request) {
	traceId := logger.GetTraceId(r)
	// firestore repositoryでのログでも使用するためcontextにtraceIdをセット
	ctx := context.WithValue(context.Background(), "traceId", traceId)
	log.Printf(logger.InfoLogEntry("[CreateUser] Application logic start", traceId))
	body, _ := ioutil.ReadAll(r.Body)

	var user *model.User
	json.Unmarshal(body, &user)
	err := ua.firestoreRepository.CreateUser(ctx, user)
	if err != nil {
		util.CreateErrorResponse(w, ctx, err, "")
		return
	}

	resModel := util.GetResponse(http.StatusCreated, "作成に成功しました。", 0, nil)
	res, _ := json.Marshal(resModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ユーザ更新
func (ua userApplication) UpdateUser(w http.ResponseWriter, r *http.Request, userId string) {
	traceId := logger.GetTraceId(r)
	// firestore repositoryでのログでも使用するためcontextにtraceIdをセット
	ctx := context.WithValue(context.Background(), "traceId", traceId)
	log.Printf(logger.InfoLogEntry("[UpdateUser] Application logic start", traceId))
	// user存在 チェック
	_, err := ua.firestoreRepository.GetUserById(ctx, userId)
	if err != nil {
		util.CreateErrorResponse(w, ctx, err, userId)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)

	var newUser *model.User
	json.Unmarshal(body, &newUser)
	newUser.Id = userId
	err = ua.firestoreRepository.UpdateUser(ctx, newUser)
	if err != nil {
		util.CreateErrorResponse(w, ctx, err, userId)
		return
	}

	resModel := util.GetResponse(http.StatusNoContent, "更新に成功しました。", 0, nil)
	res, _ := json.Marshal(resModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// ユーザ削除
func (ua userApplication) DeleteUser(w http.ResponseWriter, r *http.Request, userId string) {
	traceId := logger.GetTraceId(r)
	// firestore repositoryでのログでも使用するためcontextにtraceIdをセット
	ctx := context.WithValue(context.Background(), "traceId", traceId)
	log.Printf(logger.InfoLogEntry("[DeleteUser] Application logic start", traceId))
	// user存在 チェック
	_, err := ua.firestoreRepository.GetUserById(ctx, userId)
	if err != nil {
		util.CreateErrorResponse(w, ctx, err, userId)
		return
	}

	err = ua.firestoreRepository.DeleteUser(ctx, userId)
	if err != nil {
		util.CreateErrorResponse(w, ctx, err, userId)
		return
	}

	resModel := util.GetResponse(http.StatusNoContent, "削除に成功しました。", 0, nil)
	res, _ := json.Marshal(resModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
