package routes

import (
	"slices"

	"github.com/csvitor-dev/go.dev/types"
)

func All() []types.Route {
	return slices.Concat(authentication, user, follower, publications)
}
