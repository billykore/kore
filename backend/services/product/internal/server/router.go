package server

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/product/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg                    *config.Config
	log                    *log.Logger
	router                 *echo.Echo
	productHandler         *handler.ProductHandler
	productCategoryHandler *handler.ProductCategoryHandler
	discountHandler        *handler.DiscountHandler
	cartHandler            *handler.CartHandler
}

func NewRouter(
	cfg *config.Config,
	log *log.Logger,
	router *echo.Echo,
	productHandler *handler.ProductHandler,
	productCategoryHandler *handler.ProductCategoryHandler,
	discountHandler *handler.DiscountHandler,
	cartHandler *handler.CartHandler,
) *Router {
	return &Router{
		cfg:                    cfg,
		log:                    log,
		router:                 router,
		productHandler:         productHandler,
		productCategoryHandler: productCategoryHandler,
		discountHandler:        discountHandler,
		cartHandler:            cartHandler,
	}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	r.router.GET("/products", r.productHandler.GetProductList)
	r.router.GET("/products/:productId", r.productHandler.GetProductById)
	r.router.GET("/categories", r.productCategoryHandler.GetCategoryList)
	r.router.GET("/discounts", r.discountHandler.GetDiscountList)
	r.router.GET("/carts", r.cartHandler.GetCartItemList)
	r.router.POST("/carts", r.cartHandler.AddCartItem)
	r.router.PUT("/carts/:cartId", r.cartHandler.UpdateCartItemQuantity)
	r.router.DELETE("/carts/:cartId", r.cartHandler.DeleteCartItem)
}

func (r *Router) useMiddlewares() {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())
}

func (r *Router) run() {
	port := r.cfg.HTTPPort
	r.log.Infof("running on port [::%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [::%v]", port)
	}
}
