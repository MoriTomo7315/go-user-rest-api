package gcplogger

import (
	"github.com/MoriTomo7315/go-user-rest-api/domain/model"
)

// INFOレベルのログ出力
func InfoLogEntry(message string) string {
	entry := &model.LogEntry{
		Severity: model.INFO,
		Message:  message,
	}

	return entry.String()
}

// WARNレベルのログ出力
func WarnLogEntry(message string) string {
	entry := &model.LogEntry{
		Severity: model.WARN,
		Message:  message,
	}

	return entry.String()
}

// ERRORレベルのログ出力
func ErrorLogEntry(message string) string {
	entry := &model.LogEntry{
		Severity: model.ERROR,
		Message:  message,
	}

	return entry.String()
}
