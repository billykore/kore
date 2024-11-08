package server

import (
	handler2 "github.com/billykore/kore/backend/infra/http/handler"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server to run.
type Server struct {
	router *Router
}

// New creates new Server.
func New(router *Router) *Server {
	return &Server{
		router: router,
	}
}

// Serve start the Server.
func (s *Server) Serve() {
	s.router.Run()
}

// Router get all request to handlers and returns the response produce by handlers.
type Router struct {
	cfg             *config.Config
	log             *logger.Logger
	router          *echo.Echo
	authHandler     *handler2.UserHandler
	orderHandler    *handler2.OrderHandler
	otpHandler      *handler2.OtpHandler
	productHandler  *handler2.ProductHandler
	shippingHandler *handler2.ShippingHandler
}

// NewRouter returns new Router.
func NewRouter(
	cfg *config.Config,
	log *logger.Logger,
	router *echo.Echo,
	authHandler *handler2.UserHandler,
	orderHandler *handler2.OrderHandler,
	otpHandler *handler2.OtpHandler,
	productHandler *handler2.ProductHandler,
	shippingHandler *handler2.ShippingHandler,
) *Router {
	return &Router{
		cfg:             cfg,
		log:             log,
		router:          router,
		authHandler:     authHandler,
		orderHandler:    orderHandler,
		otpHandler:      otpHandler,
		productHandler:  productHandler,
		shippingHandler: shippingHandler,
	}
}

func (r *Router) useMiddlewares() {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())
}

func (r *Router) run() {
	port := r.cfg.HTTPPort
	r.log.Infof("running on port ::[:%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [::%v]", port)
	}
}

func (r *Router) Run() {
	r.setLoginRoutes()
	r.setProductRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setLoginRoutes() {
	r.router.POST("/login", r.authHandler.Login)
	r.router.POST("/logout", r.authHandler.Logout)
}

func (r *Router) setProductRoutes() {
	r.setCartRoutes()

	r.router.GET("/products", r.productHandler.GetProductList)
	r.router.GET("/products/:productId", r.productHandler.GetProductById)
	r.router.GET("/categories", r.productHandler.GetCategoryList)
	r.router.GET("/discounts", r.productHandler.GetDiscountList)
}

func (r *Router) setCartRoutes() {
	cr := r.router.Group("/carts")
	cr.Use(AuthMiddleware())

	cr.GET("", r.productHandler.GetCartItemList)
	cr.POST("", r.productHandler.AddCartItem)
	cr.PUT("/:cartId", r.productHandler.UpdateCartItemQuantity)
	cr.DELETE("/:cartId", r.productHandler.DeleteCartItem)
}
