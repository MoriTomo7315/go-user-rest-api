package error

import "errors"

var (
	SYSTEM_ERR     = errors.New("system error occured.")
	NOT_FOUND_USER = errors.New("user is not found")

	SYSTEM_ERR_MSG         = "システムエラーが発生しました。"
	NOT_FOUND_USER_ERR_MSG = "指定のユーザが見つかりませんでした。"
)
