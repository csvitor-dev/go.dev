package types

import (
	"net/http"
	"slices"
)

type MiddlewarePipeline struct {
	middlewares []Middleware
}

func NewPipeline(queue ...Middleware) *MiddlewarePipeline {
	slices.Reverse(queue)

	return &MiddlewarePipeline{
		middlewares: queue,
	}
}

func (m *MiddlewarePipeline) Apply(target http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range m.middlewares {
		target = middleware(target)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		target(w, r)
	}
}
