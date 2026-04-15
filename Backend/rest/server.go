package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
)

type Server struct {
	config         *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	config *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		config:         config,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()

	manager.LoadMiddleware(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// Initialize routes
	// initRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	port := ":" + server.config.HttpPort
	fmt.Println("Server Running on:", port)

	//try to catch error
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
