package routes

import (
	"net/http"

	users "github.com/csvitor-dev/social-media/src/api/controllers"
	"github.com/csvitor-dev/social-media/src/api/middlewares"
)

var userRoutes = []Route{
	{
		Uri:                "/users",
		Method:             http.MethodGet,
		Handler:            users.GetAllUsers,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodGet,
		Handler:            users.GetUserById,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/users",
		Method:             http.MethodPost,
		Handler:            users.CreateUser,
		MiddlewarePipeline: middlewares.SignPipeline(),
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodPut,
		Handler:            users.UpdateUserById,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodDelete,
		Handler:            users.DeleteUserById,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
}
