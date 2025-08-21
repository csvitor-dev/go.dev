package types

import (
	"net/http"
	"slices"
)

type MiddlewarePipeline struct {
	Middlewares []Middleware
}

func NewPipeline() *MiddlewarePipeline {
	return &MiddlewarePipeline{
		Middlewares: []Middleware{},
	}
}

func (m *MiddlewarePipeline) With(queue ...Middleware) *MiddlewarePipeline {
	slices.Reverse(queue)
	m.Middlewares = queue

	return m
}

func (m *MiddlewarePipeline) All() []Middleware {
	return m.Middlewares
}

func (m *MiddlewarePipeline) Apply(target http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range m.Middlewares {
		target = middleware(target)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		target(w, r)
	}
}
