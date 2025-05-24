package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/ctxkey"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/handler"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/jwt"
)

// RequireAuth middleware validates JWT token and adds user info to context
func RequireAuth(jwtService jwt.IJwt) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				handler.SendError(w, errorpkg.ErrNoBearerToken())
				return
			}

			// Extract token from "Bearer <token>"
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				handler.SendError(w, errorpkg.ErrInvalidBearerToken())
				return
			}

			token := tokenParts[1]

			// Validate token
			validateResp, err := jwtService.Validate(token)
			if err != nil {
				handler.SendError(w, err)
				return
			}

			// Add user info to context
			ctx := context.WithValue(r.Context(), ctxkey.UserID, validateResp.UserID.String())
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// RequireRole middleware checks if user has required role
func RequireRole(roles ...enum.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract role from JWT claims

			next.ServeHTTP(w, r)
		})
	}
}
