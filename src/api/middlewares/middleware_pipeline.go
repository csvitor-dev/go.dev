package middlewares

type MiddlewarePipeline struct {
	middlewares []Middleware
}

func SignPipeline() *MiddlewarePipeline {
	return &MiddlewarePipeline{
		middlewares: []Middleware{},
	}
}

func (m *MiddlewarePipeline) All() []Middleware {
	return m.middlewares
}
