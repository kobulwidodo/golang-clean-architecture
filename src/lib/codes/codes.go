package codes

import "net/http"

type Message struct {
	HttpCode int
	Message  string
}

var (
	// 4xx
	CodeUnauthorized = Message{
		HttpCode: http.StatusUnauthorized,
		Message:  "Unauthorized access",
	}

	// SQL
	CodeSQLRecordDoesNotExist = Message{
		HttpCode: http.StatusNotFound,
		Message:  "Record does not exist",
	}
	CodeSQLRead = Message{
		HttpCode: http.StatusInternalServerError,
		Message:  "Internal Server Error",
	}

	CodeSuccess = Message{
		HttpCode: http.StatusOK,
		Message:  "Request successful",
	}
)
