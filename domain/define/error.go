package define

import "errors"

var (
	SYSTEM_ERR     = errors.New("system error occured.")
	NOT_FOUND_USER = errors.New("user is not found")
	FAILED_CREATE_USER = errors.New("failed to create user")
	FAILED_UPDATE_USER = errors.New("failed to update user")
	FAILED_DELETE_USER = errors.New("failed to delete user")

	SYSTEM_ERR_MSG         = "システムエラーが発生しました。"
	NOT_FOUND_USER_ERR_MSG = "指定のユーザが見つかりませんでした。"
	FAILED_CREATE_USER_MSG = "ユーザの作成に失敗しました"
	FAILED_UPDATE_USER_MSG = "ユーザの更新に失敗しました"
	FAILED_DELETE_USER_MSG = "ユーザの削除に失敗しました"
)
