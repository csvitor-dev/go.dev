package routes

import (
	"slices"

	"github.com/csvitor-dev/social-media/types"
)

func All() []types.Route {
	return slices.Concat(authentication, user, follower)
}
