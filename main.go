package main

import (
	"log"
	"net/http"

	"github.com/MoriTomo7315/go-user-rest-api/application"
	"github.com/MoriTomo7315/go-user-rest-api/controller"
	"github.com/MoriTomo7315/go-user-rest-api/infrastructure/persistence"
)

func main() {
	//log設定
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

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
	log.Printf("/api/users   start")

	http.HandleFunc("/api/users", userController.HandlerHttpRequest)
	http.HandleFunc("/api/users/", userController.HandlerHttpRequestWithParameter)

	server.ListenAndServe()
}
