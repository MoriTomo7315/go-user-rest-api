package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MoriTomo7315/go-user-rest-api/application"
	"github.com/MoriTomo7315/go-user-rest-api/controller"
	logger "github.com/MoriTomo7315/go-user-rest-api/gcplogger"
	"github.com/MoriTomo7315/go-user-rest-api/infrastructure/persistence"
	"github.com/joho/godotenv"
)

func main() {

	//log設定
	log.SetFlags(0)
	// log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	// envファイル読み込み
	_ = godotenv.Load(fmt.Sprintf("./.env.%s", os.Getenv("GO_ENV")))

	// Start listening port
	server := http.Server{
		Addr: ":50001",
	}

	/*
		DDD依存関係を定義
		Infrastructure → Application → Controller
	*/
	// firestore用infrastructure
	firestoreClient := persistence.NewFirestoreClient()
	userApplication := application.NewUserApplication(firestoreClient)
	userController := controller.NewUserController(userApplication)
	//サーバーにController(ハンドラ)を登録
	log.Printf(logger.InfoLogEntry("/api/users   start", ""))

	http.HandleFunc("/api/users", userController.HandlerHttpRequest)
	http.HandleFunc("/api/users/", userController.HandlerHttpRequestWithParameter)

	server.ListenAndServe()
}
