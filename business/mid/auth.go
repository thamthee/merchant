package mid

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	"github.com/thamthee/merchant/business/auth"
)

func Authenticate(a *auth.Auth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authStr := r.Header.Get("authorization")

			parts := strings.Split(authStr, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				http.Error(w, "expected authorization header format: bearer <token>", http.StatusUnauthorized)
				return
			}

			claims, err := a.ValidateToken(parts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), auth.Key, claims)

			r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func Authorize(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(auth.Key).(auth.Claims)
			if !ok {
				http.Error(w, "claims missing from context", http.StatusUnauthorized)
				return
			}

			if !claims.Authorized(roles...) {
				http.Error(
					w,
					fmt.Errorf("you are not authorized for that action: claims: %v exp: %v", claims.Roles, roles).Error(),
					http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func BypassToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authStr := r.Header.Get("authorization")

		parts := strings.Split(authStr, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "expected authorization header format: bearer <token>", http.StatusUnauthorized)
			return
		}

		claims := auth.Claims{
			StandardClaims: jwt.StandardClaims{
				Issuer:    "service project",
				ExpiresAt: jwt.At(time.Now().Add(time.Hour)),
				IssuedAt:  jwt.At(time.Now()),
				Subject:   parts[1],
			},
			Roles: []string{auth.RoleAdmin},
		}

		ctx := context.WithValue(r.Context(), auth.Key, claims)

		r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
