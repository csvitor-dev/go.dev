package routes

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/src/apps/web/controllers/views"
	"github.com/csvitor-dev/go.dev/src/middlewares"
	"github.com/csvitor-dev/go.dev/types"
)

var guardViews = []types.Route{
	{
		Uri:                "/home",
		Method:             http.MethodGet,
		Handler:            views.GetHomeView,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger, middlewares.RequireAuth),
	},
}
