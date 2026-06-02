package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareManager struct {
	middlewares []Middleware
}

func NewMiddlewareManager() *MiddlewareManager {
	middlewaresData := MiddlewareManager{
		middlewares: make([]Middleware, 0),
	}

	return &middlewaresData
} 

func (m *MiddlewareManager) Use(middleware ...Middleware) {
	m.middlewares = append(m.middlewares, middleware...)
}

func (m *MiddlewareManager) AddMiddleware(next http.Handler, middleware ...Middleware) http.Handler {
	for _, mid := range middleware {
		next = mid(next)
	}
	return next
}

func (m *MiddlewareManager) Apply(next http.Handler) http.Handler {
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		next = m.middlewares[i](next)
	}
	return next
}
