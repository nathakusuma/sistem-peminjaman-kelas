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
			ctx := context.WithValue(r.Context(), ctxkey.UserEmail, validateResp.UserEmail)
			ctx = context.WithValue(ctx, ctxkey.UserRole, validateResp.Role)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// RequireRole middleware checks if user has required role
func RequireRole(role enum.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value(ctxkey.UserRole).(enum.UserRole)
			if !ok {
				handler.SendError(w, errorpkg.ErrForbiddenRole())
				return
			}
			if userRole != role {
				handler.SendError(w, errorpkg.ErrForbiddenRole())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
