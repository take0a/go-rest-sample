package utils

import (
	"fmt"
	"log"
	"net/http"
)

// ErrorReturn は、エラー発生後、中断する場合の共通処理
func ErrorReturn(w http.ResponseWriter, statusCode int, err error, msg string) {
	str := fmt.Sprintf("%s %s", msg, err)
	log.Println(str)
	w.WriteHeader(statusCode)
	w.Write([]byte(str))
}

// ServerError は、サーバー起因のエラー発生後、中断する場合の処理
func ServerError(w http.ResponseWriter, err error, msg string) {
	ErrorReturn(w, http.StatusInternalServerError, err, msg)
}

// BadRequest は、クライアント起因のエラー発生後、中断する場合の処理
func BadRequest(w http.ResponseWriter, err error, msg string) {
	ErrorReturn(w, http.StatusBadRequest, err, msg)
}

// NotFound は、対象が存在しないため、中断する場合の処理
func NotFound(w http.ResponseWriter, err error, msg string) {
	ErrorReturn(w, http.StatusNotFound, err, msg)
}

// Conflict は、対象が重複したため、中断する場合の処理
func Conflict(w http.ResponseWriter, err error, msg string) {
	ErrorReturn(w, http.StatusConflict, err, msg)
}
