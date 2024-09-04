package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/security/token"
)

func extractToken(req *http.Request) (string, error) {
	headerToken := req.Header.Get("Authorization")
	if headerToken == "" {
		return "", errors.New("no authorization header")
	}
	tokenString := strings.Split(headerToken, "Bearer ")
	if len(tokenString) != 2 {
		return "", errors.New("invalid authorization header")
	}
	return tokenString[1], nil
}

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		authToken, err := extractToken(req)
		if err != nil {
			code, res := entity.ResponseUnauthorized(err)
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(code)
			err := json.NewEncoder(rw).Encode(res)
			if err != nil {
				return
			}
			return
		}

		user, err := token.Verify(authToken)
		if err != nil {
			code, res := entity.ResponseUnauthorized(err)
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(code)
			err := json.NewEncoder(rw).Encode(res)
			if err != nil {
				return
			}
			return
		}

		ctx := req.Context()
		ctx = ctxt.ContextWithUser(ctx, user)
		req = req.WithContext(ctx)

		h.ServeHTTP(rw, req)
	})
}
