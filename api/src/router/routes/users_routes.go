package routes

import (
	"net/http"

	c "github.com/csvitor-dev/social-media/api/src/controllers"
)

var baseUri = "/users"
var paramId = baseUri + "/{id}"

var getAllUsers = Route{
	URI: baseUri,
	Method: http.MethodGet,
	Handler: c.GetAllUsers,
	RequireAuth: false,
}
var getUserByID = Route{
	URI: paramId,
	Method: http.MethodGet,
	Handler: c.GetUserByID,
	RequireAuth: false,
}
var createUser = Route{
	URI: baseUri,
	Method: http.MethodPost,
	Handler: c.CreateUser,
	RequireAuth: false,
}
var updateUser = Route{
	URI: paramId,
	Method: http.MethodPut,
	Handler: c.UpdateUserByID,
	RequireAuth: false,
}
var deleteUser = Route{
	URI: paramId,
	Method: http.MethodDelete,
	Handler: c.DeleteUserByID,
	RequireAuth: false,
}

var UserRoutes = []Route{
	getAllUsers, getUserByID,
	createUser, updateUser,
	deleteUser,
}