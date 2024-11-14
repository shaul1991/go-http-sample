package middleware

import (
	"encoding/json"
	"net/http"

	"go-http/exceptions"
)

// ErrorHandler is a middleware that handles errors
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				var appErr *exceptions.AppError
				switch t := rec.(type) {
				case *exceptions.AppError:
					appErr = t
				case error:
					appErr = exceptions.InternalServerError(t)
				default:
					appErr = exceptions.InternalServerError(nil)
				}

				// Check if the request expects a JSON response
				if r.Header.Get("Content-Type") == "application/json" {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(appErr.Code)
					json.NewEncoder(w).Encode(map[string]interface{}{
						"code":    appErr.Code,
						"message": appErr.Message,
					})
				} else {
					// Default to plain text response
					w.WriteHeader(appErr.Code)
					w.Write([]byte(appErr.Message))
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
