package middleware

import "net/http"

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

	for _, middleware := range middleware {
		next = middleware(next)
	}

	for _, globalMiddleware := range mngr.globalMiddlewares {
		next = globalMiddleware(next)
	}

	return next
}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {

	for _, globalMiddleware := range mngr.globalMiddlewares {
		next = globalMiddleware(next)
	}

	return next
}
