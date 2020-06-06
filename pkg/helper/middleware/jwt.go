package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/yoanyombapro/CubeMicroservices/podinfo/pkg/utils/jwtutils"
)

// JwtMiddleware witholds jwt issuer and secret
type JwtMiddleware struct {
	Issuer string
	Secret string
}

// NewJwtMiddleware returns a new pointer reference to a JwtMiddleware
func NewJwtMiddleware(issuer, secret string) *JwtMiddleware {
	return &JwtMiddleware{
		Issuer: issuer,
		Secret: secret,
	}
}

// Handler is the jwt middleware handler function. It extracts a jwt token
// from a request header and asserts it is still valid in a middleware chain
func (m *JwtMiddleware) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Malformed Token"))
		}
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("authorization bearer header required"))
		}

		claims := jwtutils.JwtCustomClaims{}
		token, err := jwt.ParseWithClaims(bearerToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
			}
			return []byte(m.Secret), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized. Invalid Token"))
			return
		}

		if token.Valid {
			if claims.StandardClaims.Issuer != m.Issuer {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("Unauthorized. Unrecognized Issuer"))
			} else {
				ctx := context.WithValue(r.Context(), "jwt-token", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("properties").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized"))
		}
	}
}

