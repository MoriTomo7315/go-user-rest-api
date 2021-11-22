package model

// エラー時のレスポンスモデルの構造体
type ErrorResponse struct {
	Code    int    `json:code`
	Message string `json:message`
}
