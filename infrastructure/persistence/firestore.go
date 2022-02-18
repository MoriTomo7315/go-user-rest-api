package persistence

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/MoriTomo7315/go-user-rest-api/domain/define"
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
	"github.com/MoriTomo7315/go-user-rest-api/infrastructure/logger"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type firestoreClient struct{}

func NewFirestoreClient() repository.FirestoreRepository {
	return &firestoreClient{}
}

func initFireStoreClient(ctx context.Context) (*firestore.Client, error) {
	// .envのGOOGLE_APPLICATION_CREDENTIALSから暗黙的に設定を読み取る
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firebase.NewAppに失敗 %s", err.Error())))
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestore client 初期化に失敗 %s", err.Error())))
		return nil, err
	}
	return client, nil
}

// firestoreから全ユーザの情報を取得する
func (f *firestoreClient) GetUsers() (users []*model.User, err error) {
	log.Printf(logger.InfoLogEntry("[GetUsers] connecting firestore start."))

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestore clientの初期化に失敗 err=%s", err.Error())))
		return nil, define.SYSTEM_ERR
	}

	iter := client.Collection("users").Documents(ctx)
	for {
		userDocSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestoreからusersコレクションの検索に失敗 err=%s", err.Error())))
			return nil, define.NOT_FOUND_USER
		}
		// Uidをmap[string]interface{}に含める
		userData := userDocSnap.Data()
		userData["id"] = userDocSnap.Ref.ID
		// map[string]interface{} →json []byte -> *model.BookingModel
		jsonuserData, _ := json.Marshal(userData)
		var user *model.User
		json.Unmarshal(jsonuserData, &user)
		users = append(users, user)
	}

	log.Printf(logger.InfoLogEntry("[GetUsers] connecting firestore end."))
	return users, nil
}

// firestoreからユーザの情報を取得する
func (f *firestoreClient) GetUserById(id string) (user *model.User, err error) {
	log.Printf(logger.InfoLogEntry(fmt.Sprint("[GetUserById] connecting firestore start. id=%s", id)))

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestore clientの初期化に失敗 err=%s", err.Error())))
		return nil, define.SYSTEM_ERR
	}

	userDocSnap, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestoreからusersコレクションの検索に失敗 id=%s, err=%s", id, err.Error())))
		return nil, define.NOT_FOUND_USER
	}

	// userドキュメントの中身を返却
	// Uidをmap[string]interface{}に含める
	userData := userDocSnap.Data()
	userData["id"] = userDocSnap.Ref.ID
	// map[string]interface{} →json []byte -> *model.BookingModel
	jsonuserData, _ := json.Marshal(userData)
	json.Unmarshal(jsonuserData, &user)

	log.Printf(logger.InfoLogEntry(fmt.Sprint("[GetUserById] connecting firestore end. id=%s", id)))
	return user, nil
}

// firestoreにユーザの情報を作成する
func (f *firestoreClient) CreateUser(user *model.User) (err error) {
	log.Printf(logger.InfoLogEntry(fmt.Sprint("[CreateUser] connecting firestore start. name=%v", user)))

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestore clientの初期化に失敗 err=%s", err.Error())))
		return define.SYSTEM_ERR
	}

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"name":       user.Name,
		"prefecture": user.Prefecture,
		"createdAt":  time.Now(),
		"updatedAt":  nil,
	})
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestoreのusersコレクションの作成に失敗 err=%s", err.Error())))
		return define.FAILED_CREATE_USER
	}

	log.Printf(logger.InfoLogEntry(fmt.Sprint("[CreateUser] connecting firestore end.")))
	return nil
}

// firestoreのユーザ情報を更新する
func (f *firestoreClient) UpdateUser(user *model.User) (err error) {
	log.Printf(logger.InfoLogEntry(fmt.Sprint("[UpdateUser] connecting firestore start. user=%v", user)))

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestore clientの初期化に失敗 err=%s", err.Error())))
		return define.SYSTEM_ERR
	}

	_, err = client.Collection("users").Doc(user.Id).Set(ctx, map[string]interface{}{
		"name":       user.Name,
		"prefecture": user.Prefecture,
		"createdAt":  user.CreatedAt,
		"updatedAt":  time.Now(),
	})

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestoreのusersコレクションの更新に失敗 id=%s, err=%s", user.Id, err.Error())))
		return define.FAILED_UPDATE_USER
	}

	log.Printf(logger.InfoLogEntry("[UpdateUser] connecting firestore end."))
	return nil
}

// firestoreのユーザ情報を削除する
func (f *firestoreClient) DeleteUser(id string) (err error) {
	log.Printf(logger.InfoLogEntry(fmt.Sprint("[DeleteUser] connecting firestore start. userId=%v", id)))

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestore clientの初期化に失敗 err=%s", err.Error())))
		return define.SYSTEM_ERR
	}

	_, err = client.Collection("users").Doc(id).Delete(ctx)

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("firestoreのusersコレクションの削除に失敗 id=%s, err=%s", id, err.Error())))
		return define.FAILED_DELETE_USER
	}

	log.Printf(logger.InfoLogEntry("[DeleteUser] connecting firestore end."))
	return nil
}
