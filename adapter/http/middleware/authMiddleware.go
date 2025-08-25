package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pluvia/pluvia-api/core/domain"
)

type AuthMiddleware struct {
	UseCase domain.AuthUseCase
}

func NewAuthMiddleware(useCase domain.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		UseCase: useCase,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenCookie, err := r.Cookie("access_token")
		if err != nil {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		token, err := jwt.ParseWithClaims(accessTokenCookie.Value, &domain.JwtToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Substitua pela sua chave secreta
		})

		if err != nil || !token.Valid {
			// Se o token de acesso for inválido ou expirado, tente usar o refresh token
			refreshTokenCookie, err := r.Cookie("refresh_token")
			if err != nil {
				http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
				return
			}

			newTokens, err := m.UseCase.RefreshToken(refreshTokenCookie.Value)
			if err != nil {
				http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
				return
			}

			// Defina os novos cookies
			http.SetCookie(w, &http.Cookie{
				Name:     "access_token",
				Value:    newTokens.Token,
				Path:     "/",
				Expires:  time.Now().Add(15 * time.Minute),
				HttpOnly: true,
			})
			http.SetCookie(w, &http.Cookie{
				Name:     "refresh_token",
				Value:    newTokens.Refresh,
				Path:     "/",
				Expires:  time.Now().Add(7 * 24 * time.Hour),
				HttpOnly: true,
			})

			// Continue com a requisição original
			next.ServeHTTP(w, r)
			return
		}

		// Se o token de acesso for válido
		next.ServeHTTP(w, r)
	})
}