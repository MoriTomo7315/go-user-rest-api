package persistence

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	myError "github.com/MoriTomo7315/go-user-rest-api/domain/error"
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
	"github.com/MoriTomo7315/go-user-rest-api/domain/repository"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type firestoreClient struct{}

func NewFirestoreClient() repository.FirestoreRepository {
	return &firestoreClient{}
}

func loadEnvFile() error {
	return godotenv.Load(fmt.Sprintf("./.env.%s", os.Getenv("GO_ENV")))
}

func initFireStoreClient(ctx context.Context) (*firestore.Client, error) {
	// .envのGOOGLE_APPLICATION_CREDENTIALSから暗黙的に設定を読み取る
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Printf("ERROR firebase.NewAppに失敗 %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("ERROR firestore client 初期化に失敗 %v", err)
		return nil, err
	}
	return client, nil
}

// firestoreから全ユーザの情報を取得する
func (f *firestoreClient) GetUsers() (users []*model.User, err error) {
	log.Printf("INFO [GetUsers] connecting firestore start.")

	err = loadEnvFile()
	if err != nil {
		// .env読めなかった場合の処理
		log.Printf("ERROR .envファイル読み込み失敗 err=%v", err)
		return nil, myError.SYSTEM_ERR
	}

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf("ERROR firestore clientの初期化に失敗 err=%v", err)
		return nil, err
	}

	iter := client.Collection("users").Documents(ctx)
	for {
		userDocSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
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

	log.Printf("INFO [GetUsers] connecting firestore end.")
	return users, nil
}

// firestoreからユーザの情報を取得する
func (f *firestoreClient) GetUserById(id string) (user *model.User, err error) {
	log.Printf("INFO [GetUserById] connecting firestore start. id=%s", id)

	err = loadEnvFile()
	if err != nil {
		// .env読めなかった場合の処理
		log.Printf("ERROR .envファイル読み込み失敗 err=%v", err)
		return nil, myError.SYSTEM_ERR
	}

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf("ERROR firestore clientの初期化に失敗 err=%v", err)
		return nil, myError.SYSTEM_ERR
	}

	userDocSnap, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("ERROR firestoreからusersコレクションの検索に失敗 id=%s, err=%v", id, err)
		return nil, myError.NOT_FOUND_USER
	}

	// userドキュメントの中身を返却
	// Uidをmap[string]interface{}に含める
	userData := userDocSnap.Data()
	userData["id"] = userDocSnap.Ref.ID
	// map[string]interface{} →json []byte -> *model.BookingModel
	jsonuserData, _ := json.Marshal(userData)
	json.Unmarshal(jsonuserData, &user)

	log.Printf("INFO [GetUserById] connecting firestore end. id=%s", id)
	return user, nil
}

// firestoreにユーザの情報を作成する
func (f *firestoreClient) CreateUser(user *model.User) (id string, err error) {
	log.Printf("INFO [CreateUser] connecting firestore start. name=%v", user)

	err = loadEnvFile()
	if err != nil {
		// .env読めなかった場合の処理
		log.Printf("ERROR .envファイル読み込み失敗 err=%v", err)
		return "", myError.SYSTEM_ERR
	}

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf("ERROR firestore clientの初期化に失敗 err=%v", err)
		return "", myError.SYSTEM_ERR
	}

	userDocRef, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"name":       user.Name,
		"prefecture": user.Prefecture,
		"createdAt":  user.CreatedAt,
		"updated":    user.UpdatedAt,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	id = userDocRef.ID

	log.Printf("INFO [CreateUser] connecting firestore end.")
	return id, nil
}

// firestoreのユーザ情報を更新する
func (f *firestoreClient) UpdateUser(user *model.User) (err error) {
	log.Printf("INFO [UpdateUser] connecting firestore start. user=%v", user)

	err = loadEnvFile()
	if err != nil {
		// .env読めなかった場合の処理
		log.Printf("ERROR .envファイル読み込み失敗 err=%v", err)
		return myError.SYSTEM_ERR
	}

	// init firestore client
	ctx := context.Background()
	client, err := initFireStoreClient(ctx)
	defer client.Close()

	if err != nil {
		log.Printf("ERROR firestore clientの初期化に失敗 err=%v", err)
		return myError.SYSTEM_ERR
	}

	_, err = client.Collection("users").Doc(user.Id).Set(ctx, map[string]interface{}{
		"name":       user.Name,
		"prefecture": user.Prefecture,
		"createdAt":  user.CreatedAt,
		"updatedAt":  time.Now(),
	})

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	log.Printf("INFO [UpdateUser] connecting firestore end.")
	return nil
}
