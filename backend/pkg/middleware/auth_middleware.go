package middleware

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/security/token"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Auth returns middleware function that validate token from headers
// and extract user information.
func Auth() echo.MiddlewareFunc {
	cfg := config.Get()
	return echojwt.WithConfig(jwtConfig(cfg))
}

// jwtConfig contains configuration for auth middleware.
func jwtConfig(cfg *config.Config) echojwt.Config {
	return echojwt.Config{
		ContextKey:     ctxt.UserContextKey,
		SigningKey:     []byte(cfg.Token.Secret),
		SuccessHandler: successHandler,
		ErrorHandler:   errorHandler,
	}
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
