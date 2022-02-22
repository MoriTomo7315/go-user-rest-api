package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MoriTomo7315/go-user-rest-api/domain/define"
	logger "github.com/MoriTomo7315/go-user-rest-api/gcplogger"
)

func CreateErrorResponse(w http.ResponseWriter, ctx context.Context, err error, userId string) {
	traceId := ctx.Value("traceId").(string)
	log.Printf(logger.InfoLogEntry("start check and return error response", traceId))
	switch err {
	case define.NOT_FOUND_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("userが見つかりませんでした。 userId=%s", userId), traceId))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		resModel := GetResponse(http.StatusNotFound, define.NOT_FOUND_USER_ERR_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_CREATE_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("userの作成に失敗しました。 userId=%s", userId), traceId))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_CREATE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_UPDATE_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("userの更新に失敗しました。 userId=%s", userId), traceId))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_UPDATE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	case define.FAILED_DELETE_USER:
		log.Printf(logger.ErrorLogEntry(fmt.Sprintf("userの削除に失敗しました。 userId=%s", userId), traceId))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.FAILED_DELETE_USER_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	default:
		log.Printf(logger.ErrorLogEntry(err.Error(), traceId))
		w.WriteHeader(http.StatusInternalServerError)
		resModel := GetResponse(http.StatusInternalServerError, define.SYSTEM_ERR_MSG, 0, nil)
		res, _ := json.Marshal(resModel)
		w.Write(res)
	}
}
