package middleware

import "net/http"

// Middleware defines the type for middleware functions that wrap HTTP handlers.
type Middleware func(http.Handler) http.HandlerFunc

// manager is responsible for managing and applying middleware to HTTP handlers.
type manager struct {
	globalMiddlewares []Middleware
}

// NewManager creates and returns a new instance of the middleware manager.
func NewManager() *manager {
	return &manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

// LoadMiddleware allows adding middleware functions to the manager's global middleware list.
func (mngr *manager) LoadMiddleware(middleware ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
}

// With applies the provided middleware to the given HTTP handler, along with any global middleware loaded in the manager.
func (mngr *manager) With(next http.Handler, middleware ...Middleware) http.Handler {

	// Apply the provided middleware in the order they were given
	for _, middleware := range middleware {
		next = middleware(next)
	}

	// Apply global middleware in the order they were loaded
	for _, globalMiddleware := range mngr.globalMiddlewares {
		next = globalMiddleware(next)
	}

	// Return the final handler wrapped with all middleware
	return next
}
