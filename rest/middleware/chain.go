package middleware

import(
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

func (m *Manager) Use(md ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, md...)
}

func (m *Manager) Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	
	// Apply route-specific middlewares first
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}

func (m *Manager) WrappedMux(h http.Handler) http.Handler {

	// Apply global middlewares after route-specific ones
	for i := len(m.globalMiddlewares) - 1; i >= 0; i-- {
		h = m.globalMiddlewares[i](h)
	}

	return h
}