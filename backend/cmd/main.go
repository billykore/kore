package main

import (
	"github.com/billykore/kore/backend/infra/http/server"
	"github.com/billykore/kore/backend/infra/messaging"
	"github.com/billykore/kore/backend/pkg/config"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	ss *server.Server
	mc *messaging.Consumer
}

func newApp(ss *server.Server, mc *messaging.Consumer) *app {
	return &app{
		ss: ss,
		mc: mc,
	}
}

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
//	@host			api.kore.co.id
//	@schemes		http https
//	@BasePath		/api/v1
func main() {
	c := config.Get()
	a := initApp(c)
	go a.mc.Consume()
	a.ss.Serve()
}
