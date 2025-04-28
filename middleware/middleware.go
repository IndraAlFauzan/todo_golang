package middleware

import (
	"belajar_go/handler"
	"net/http"
)

// MethodNotAllowedHandler untuk menangani 405 error dengan response JSON
func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.WriteJSONResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
	})
}
