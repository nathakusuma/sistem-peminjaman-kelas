package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/config"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get frontend URL from config
		frontendURL := config.GetEnv().FrontendURL
		if frontendURL == "" {
			// replace https with ""
			frontendURL = strings.Replace(frontendURL, "https://", "", 1)
			fmt.Println(frontendURL)
		}

		w.Header().Set("Access-Control-Allow-Origin", frontendURL)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
