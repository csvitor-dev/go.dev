package routes

import (
	"net/http"

	"github.com/csvitor-dev/social-media/src/api/controllers"
	"github.com/csvitor-dev/social-media/src/api/middlewares"
)

var followers = []Route{
	{
		Uri:                "/follow/{userId}",
		Method:             http.MethodPost,
		handler:            controllers.Follow,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/unfollow/{userId}",
		Method:             http.MethodPost,
		handler:            controllers.Unfollow,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/followers/{userId}",
		Method:             http.MethodGet,
		handler:            controllers.GetFollowers,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
	{
		Uri:                "/following/{userId}",
		Method:             http.MethodGet,
		handler:            controllers.GetFollowing,
		MiddlewarePipeline: middlewares.SignPipeline().AddAuthZ(),
	},
}
