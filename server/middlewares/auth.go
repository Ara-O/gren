package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session_id, err := r.Cookie("session_id")

		if err != nil {
			http.Error(w, "Error retrieving session id", http.StatusInternalServerError)
			return
		}

		fmt.Println("Session ID: ", session_id.Value)
		if strings.TrimSpace(session_id.Value) == "" {
			http.Error(w, "Error retrieving session id", http.StatusTeapot)
			return
		}

		fmt.Println(r.Cookie("session_id"))
		fmt.Println("Middleware kicks in")
		next.ServeHTTP(w, r)
	}
}
