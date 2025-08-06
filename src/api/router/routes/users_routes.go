package routes

import (
	"net/http"

	users "github.com/csvitor-dev/social-media/src/api/controllers"
)

var baseUri = "/users"
var paramId = baseUri + "/{id}"

var getAllUsers = Route{
	URI:         baseUri,
	Method:      http.MethodGet,
	Handler:     users.GetAllUsers,
	RequireAuth: false,
}
var getUserByID = Route{
	URI:         paramId,
	Method:      http.MethodGet,
	Handler:     users.GetUserByID,
	RequireAuth: false,
}
var createUser = Route{
	URI:         baseUri,
	Method:      http.MethodPost,
	Handler:     users.CreateUser,
	RequireAuth: false,
}
var updateUser = Route{
	URI:         paramId,
	Method:      http.MethodPut,
	Handler:     users.UpdateUserByID,
	RequireAuth: false,
}
var deleteUser = Route{
	URI:         paramId,
	Method:      http.MethodDelete,
	Handler:     users.DeleteUserByID,
	RequireAuth: false,
}

var UserRoutes = []Route{
	getAllUsers, getUserByID,
	createUser, updateUser,
	deleteUser,
}
