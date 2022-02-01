package util

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/MoriTomo7315/go-user-rest-api/domain/define"
)


func CreateErrorResponse(w http.ResponseWriter, err error, userId string) {
	log.Printf("INFO start check and return error response")
	switch err {
	case define.NOT_FOUND_USER:
		log.Printf("ERROR userが見つかりませんでした。 userId=%s", userId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		resModel := GetResponse(http.StatusNotFound, define.NOT_FOUND_USER_ERR_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_CREATE_USER:
		log.Printf("ERROR userの作成に失敗しました。 userId=%s", userId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_CREATE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_UPDATE_USER:
		log.Printf("ERROR userの更新に失敗しました。 userId=%s", userId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_UPDATE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_DELETE_USER:
		log.Printf("ERROR userの削除に失敗しました。 userId=%s", userId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_DELETE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	default:
		log.Printf("ERROR %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.SYSTEM_ERR_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	}
}

