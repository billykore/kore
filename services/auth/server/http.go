package server

type HTTPServer struct {
	router *Router
}

func NewHTTPServer(router *Router) *HTTPServer {
	return &HTTPServer{
		router: router,
	}
}

func (hs *HTTPServer) Serve() {
	hs.router.Run()
}
