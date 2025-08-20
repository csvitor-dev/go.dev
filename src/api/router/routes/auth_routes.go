package routes

import (
	"net/http"

	"github.com/csvitor-dev/social-media/src/api/controllers"
	"github.com/csvitor-dev/social-media/src/api/middlewares"
)

var auth = []Route{
	{
		Uri:                "/auth/register",
		Method:             http.MethodPost,
		handler:            controllers.Register,
		MiddlewarePipeline: middlewares.SignPipeline(),
	},
	{
		Uri:                "/auth/login",
		Method:             http.MethodPost,
		handler:            controllers.Login,
		MiddlewarePipeline: middlewares.SignPipeline(),
	},
}
