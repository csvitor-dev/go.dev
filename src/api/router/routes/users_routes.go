package routes

import (
	"net/http"

	users "github.com/csvitor-dev/social-media/src/api/controllers"
	"github.com/csvitor-dev/social-media/src/api/middlewares"
)

var userRoutes = []Route{
	{
		Uri:            "/users",
		Method:         http.MethodGet,
		Handler:        users.GetAllUsers,
		MiddlewareTags: middlewares.NewTags(),
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodGet,
		Handler:        users.GetUserById,
		MiddlewareTags: middlewares.NewTags(),
	},
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Handler:        users.CreateUser,
		MiddlewareTags: middlewares.NewTags(),
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodPut,
		Handler:        users.UpdateUserById,
		MiddlewareTags: middlewares.NewTags(),
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodDelete,
		Handler:        users.DeleteUserById,
		MiddlewareTags: middlewares.NewTags(),
	},
}
