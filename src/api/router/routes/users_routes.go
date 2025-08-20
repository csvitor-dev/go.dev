package routes

import (
	"net/http"

	users "github.com/csvitor-dev/social-media/src/api/controllers"
	"github.com/csvitor-dev/social-media/src/api/middlewares"
)

var user = []Route{
	{
		Uri:                "/users",
		Method:             http.MethodGet,
		handler:            users.GetAllUsers,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodGet,
		handler:            users.GetUserById,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodPut,
		handler:            users.UpdateUserById,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodDelete,
		handler:            users.DeleteUserById,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
}
