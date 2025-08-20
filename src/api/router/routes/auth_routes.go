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
	{
		Uri:                "/auth/password/refresh",
		Method:             http.MethodPost,
		handler:            controllers.RefreshPassword,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/auth/password/recover",
		Method:             http.MethodPost,
		handler:            controllers.RecoverPassword,
		MiddlewarePipeline: middlewares.SignPipeline(),
	},
	{
		Uri:                "/auth/password/reset",
		Method:             http.MethodGet,
		handler:            controllers.ValidateResetPasswordToken,
		MiddlewarePipeline: middlewares.SignPipeline(),
	},
	{
		Uri:                "/auth/password/reset",
		Method:             http.MethodPost,
		handler:            controllers.ResetPassword,
		MiddlewarePipeline: middlewares.SignPipeline(),
	},
}
