package server

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/product/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg                *config.Config
	log                *log.Logger
	router             *echo.Echo
	productSvc         *handler.ProductHandler
	productCategorySvc *handler.ProductCategoryHandler
	discountSvc        *handler.DiscountHandler
	cartSvc            *handler.CartHandler
}

func NewRouter(
	cfg *config.Config,
	log *log.Logger,
	router *echo.Echo,
	productSvc *handler.ProductHandler,
	productCategorySvc *handler.ProductCategoryHandler,
	discountSvc *handler.DiscountHandler,
	cartSvc *handler.CartHandler,
) *Router {
	return &Router{
		cfg:                cfg,
		log:                log,
		router:             router,
		productSvc:         productSvc,
		productCategorySvc: productCategorySvc,
		discountSvc:        discountSvc,
		cartSvc:            cartSvc,
	}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	r.router.GET("/products", r.productSvc.GetProductList)
	r.router.GET("/products/:productId", r.productSvc.GetProductById)
	r.router.GET("/categories", r.productCategorySvc.GetCategoryList)
	r.router.GET("/discounts", r.discountSvc.GetDiscountList)
	r.router.GET("/carts", r.cartSvc.GetCartItemList)
	r.router.POST("/carts", r.cartSvc.AddCartItem)
	r.router.PUT("/carts/:cartId", r.cartSvc.UpdateCartItemQuantity)
	r.router.DELETE("/carts/:cartId", r.cartSvc.DeleteCartItem)
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
