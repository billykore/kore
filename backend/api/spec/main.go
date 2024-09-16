package main

import (
	_ "github.com/billykore/kore/backend/api/spec/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
)

// main swaggo annotation.
//
//	@title			Gateway API
//	@version		1.0
//	@description	Gateway service API specification.
//	@termsOfService	https://swagger.io/terms/
//	@contact.name	Billy Kore
//	@contact.url	https://www.swagger.io/support
//	@contact.email	billyimmcul2010@gmail.com
//	@license.name	Apache 2.0
//	@license.url	https://www.apache.org/licenses/LICENSE-2.0.html
//	@host			https://gateway.kore.co.id
//	@BasePath		/api/v1
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/swagger/*", echoswagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
