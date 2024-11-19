package server

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/security/token"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// jwtConfig contains configuration for JWT auth middleware.
var jwtConfig = echojwt.Config{
	ContextKey:     ctxt.UserContextKey,
	SigningKey:     []byte(config.Get().Token.Secret),
	SuccessHandler: successHandler,
	ErrorHandler:   errorHandler,
}

// authMiddleware returns middleware function that validate token from headers
// and extract user information.
func authMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(jwtConfig)
}

// successHandler extract user information from token
// and save the information in the request context.
func successHandler(ctx echo.Context) {
	t := ctx.Get(ctxt.UserContextKey).(*jwt.Token)
	user := token.UserFromToken(t)
	c := ctx.Request().Context()
	c = ctxt.ContextWithUser(c, user)
	ctx.SetRequest(ctx.Request().WithContext(c))
}

// errorHandler returns unauthorized response if there is authentication error.
func errorHandler(ctx echo.Context, err error) error {
	return ctx.JSON(entity.ResponseUnauthorized(err))
}
