package persistence

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"context"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/MoriTomo7315/go-user-rest-api/domain/define"
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
	logger "github.com/MoriTomo7315/go-user-rest-api/gcplogger"
	"google.golang.org/api/iterator"
)

type firestoreClient struct{}

func NewFirestoreClient() repository.FirestoreRepository {
	return &firestoreClient{}
}

func initFireStoreClient(ctx context.Context) (*firestore.Client, error) {
	// contextからtraceIdを取得
	traceId := ctx.Value("traceId").(string)
	// .envのGOOGLE_APPLICATION_CREDENTIALSから暗黙的に設定を読み取る
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firebase.NewAppに失敗 %v", err), traceId))
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestore client 初期化に失敗 %s", err), traceId))
		return nil, err
	}
	return client, nil
}

// firestoreから全ユーザの情報を取得する
func (f *firestoreClient) GetUsers(ctx context.Context) (users []*model.User, err error) {
	// contextからtraceIdを取得
	traceId := ctx.Value("traceId").(string)
	log.Printf(logger.InfoLogEntry("[GetUsers] connecting firestore start.", traceId))

	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestore clientの初期化に失敗 err=%v", err), traceId))
		return nil, define.SYSTEM_ERR
	}

	iter := client.Collection("users").Documents(ctx)
	for {
		userDocSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestoreからusersコレクションの検索に失敗 err=%v", err), traceId))
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

	log.Printf(logger.InfoLogEntry("[GetUsers] connecting firestore end.", traceId))
	return users, nil
}

// firestoreからユーザの情報を取得する
func (f *firestoreClient) GetUserById(ctx context.Context, id string) (user *model.User, err error) {
	// contextからtraceIdを取得
	traceId := ctx.Value("traceId").(string)
	log.Printf(logger.InfoLogEntry(fmt.Sprintf("[GetUserById] connecting firestore start. id=%s", id), traceId))

	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestore clientの初期化に失敗 err=%v", err), traceId))
		return nil, define.SYSTEM_ERR
	}

	userDocSnap, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestoreからusersコレクションの検索に失敗 id=%s, err=%v", id, err), traceId))
		return nil, define.NOT_FOUND_USER
	}

	// userドキュメントの中身を返却
	// Uidをmap[string]interface{}に含める
	userData := userDocSnap.Data()
	userData["id"] = userDocSnap.Ref.ID
	// map[string]interface{} →json []byte -> *model.BookingModel
	jsonuserData, _ := json.Marshal(userData)
	json.Unmarshal(jsonuserData, &user)

	log.Printf(logger.InfoLogEntry(fmt.Sprintf("[GetUserById] connecting firestore end. id=%s", id), traceId))
	return user, nil
}

// firestoreにユーザの情報を作成する
func (f *firestoreClient) CreateUser(ctx context.Context, user *model.User) (err error) {
	// contextからtraceIdを取得
	traceId := ctx.Value("traceId").(string)
	log.Printf(logger.InfoLogEntry(fmt.Sprintf("[CreateUser] connecting firestore start. name=%v", user), traceId))

	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestore clientの初期化に失敗 err=%v", err), traceId))
		return define.SYSTEM_ERR
	}

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"name":       user.Name,
		"prefecture": user.Prefecture,
		"createdAt":  time.Now(),
		"updatedAt":  nil,
	})
	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestoreのusersコレクションの作成に失敗 err=%v", err), traceId))
		return define.FAILED_CREATE_USER
	}

	log.Printf(logger.InfoLogEntry(fmt.Sprintf("[CreateUser] connecting firestore end."), traceId))
	return nil
}

// firestoreのユーザ情報を更新する
func (f *firestoreClient) UpdateUser(ctx context.Context, user *model.User) (err error) {
	// contextからtraceIdを取得
	traceId := ctx.Value("traceId").(string)
	log.Printf(logger.InfoLogEntry(fmt.Sprintf("[UpdateUser] connecting firestore start. user=%v", user), traceId))

	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestore clientの初期化に失敗 err=%v", err), traceId))
		return define.SYSTEM_ERR
	}

	_, err = client.Collection("users").Doc(user.Id).Set(ctx, map[string]interface{}{
		"name":       user.Name,
		"prefecture": user.Prefecture,
		"createdAt":  user.CreatedAt,
		"updatedAt":  time.Now(),
	})

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestoreのusersコレクションの更新に失敗 id=%s, err=%v", user.Id, err), traceId))
		return define.FAILED_UPDATE_USER
	}

	log.Printf(logger.InfoLogEntry("[UpdateUser] connecting firestore end.", traceId))
	return nil
}

// firestoreのユーザ情報を削除する
func (f *firestoreClient) DeleteUser(ctx context.Context, id string) (err error) {
	// contextからtraceIdを取得
	traceId := ctx.Value("traceId").(string)
	log.Printf(logger.InfoLogEntry(fmt.Sprintf("[DeleteUser] connecting firestore start. userId=%v", id), traceId))

	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestore clientの初期化に失敗 err=%v", err), traceId))
		return define.SYSTEM_ERR
	}

	_, err = client.Collection("users").Doc(id).Delete(ctx)

	if err != nil {
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("firestoreのusersコレクションの削除に失敗 id=%s, err=%v", id, err), traceId))
		return define.FAILED_DELETE_USER
	}

	log.Printf(logger.InfoLogEntry("[DeleteUser] connecting firestore end.", traceId))
	return nil
}
