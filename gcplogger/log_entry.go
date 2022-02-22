package gcplogger

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// ログレベルのCONSTを定義
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity
var (
	INFO  = "INFO"
	WARN  = "WARNING"
	ERROR = "ERROR"
)

// GCPのLogEntryに則った構造化ログモデル
type LogEntry struct {
	// GCP上でLogLevelを表す
	Severity string `json:"severity"`
	// ログの内容
	Message string `json:"message"`
	// トレースID
	Trace string `json:"trace"`
}

// 構造体をJSON形式の文字列へ変換
// 参考: https://cloud.google.com/run/docs/logging#run_manual_logging-go
func (l LogEntry) String() string {
	if l.Severity == "" {
		l.Severity = INFO
	}
	out, err := json.Marshal(l)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	return string(out)
}

// INFOレベルのログ出力
func InfoLogEntry(message string, trace string) string {
	entry := &LogEntry{
		Severity: INFO,
		Message:  message,
		Trace:    trace,
	}

	return entry.String()
}

// WARNレベルのログ出力
func WarnLogEntry(message string, trace string) string {
	entry := &LogEntry{
		Severity: WARN,
		Message:  message,
		Trace:    trace,
	}

	return entry.String()
}

// ERRORレベルのログ出力
func ErrorLogEntry(message string, trace string) string {
	entry := &LogEntry{
		Severity: ERROR,
		Message:  message,
		Trace:    trace,
	}

	return entry.String()
}

func GetTraceId(r *http.Request) string {
	traceHeader := r.Header.Get("X-Cloud-Trace-Context")
	traceParts := strings.Split(traceHeader, "/")
	traceId := ""
	if len(traceParts) > 0 {
		traceId = traceParts[0]
	}
	return traceId
}
