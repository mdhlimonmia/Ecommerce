package middleware

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) LoadMiddleware(middleware ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
}

func (mngr *Manager) With(next http.Handler, middleware ...Middleware) http.Handler {

	fmt.Println("............hit middleware manager...........")
	for _, middleware := range middleware {
		next = middleware(next)
	}

	return next
}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {

	fmt.Println("hit Global Middleware...........")
	for _, globalMiddleware := range mngr.globalMiddlewares {
		next = globalMiddleware(next)
	}

	return next
}
