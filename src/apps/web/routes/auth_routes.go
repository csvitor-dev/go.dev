package routes

import (
	"net/http"

	"github.com/csvitor-dev/social-media/src/apps/web/controllers"
	"github.com/csvitor-dev/social-media/types"
)

var auth = []types.Route{
	{
		Uri:                "/",
		Method:             http.MethodGet,
		Handler:            controllers.GetLoginView,
		MiddlewarePipeline: types.NewPipeline(),
	},
	{
		Uri:                "/auth/login",
		Method:             http.MethodGet,
		Handler:            controllers.GetLoginView,
		MiddlewarePipeline: types.NewPipeline(),
	},
	{
		Uri:                "/auth/reset-password",
		Method:             http.MethodGet,
		Handler:            controllers.FetchApiForTokenValidation,
		MiddlewarePipeline: types.NewPipeline(),
	},
}
