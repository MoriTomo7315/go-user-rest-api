package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MoriTomo7315/go-user-rest-api/domain/define"
	"github.com/MoriTomo7315/go-user-rest-api/infrastructure/logger"
)

func CreateErrorResponse(w http.ResponseWriter, err error, userId string) {
	log.Printf(logger.InfoLogEntry("start check and return error response"))
	switch err {
	case define.NOT_FOUND_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("userが見つかりませんでした。 userId=%s", userId)))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		resModel := GetResponse(http.StatusNotFound, define.NOT_FOUND_USER_ERR_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_CREATE_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("userの作成に失敗しました。 userId=%s", userId)))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_CREATE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_UPDATE_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("userの更新に失敗しました。 userId=%s", userId)))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_UPDATE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_DELETE_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprint("userの削除に失敗しました。 userId=%s", userId)))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_DELETE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	default:
		log.Printf(logger.ErrorLogEntry(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.SYSTEM_ERR_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	}
}
