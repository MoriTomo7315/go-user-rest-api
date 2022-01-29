package model

// エラー時のレスポンスモデルの構造体
type ErrorResponse struct {
	Message string `json:message`
}
