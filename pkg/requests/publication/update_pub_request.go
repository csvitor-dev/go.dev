package publication

import (
	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type UpdatePubRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *UpdatePubRequest) Validate() types.RequestValidationGuard {
	titleErrors := validations.NewString(r.Title, "title").IsOptional().MaxLength(100).Result()

	contentErrors := validations.NewString(r.Content, "content").IsOptional().MaxLength(300).Result()

	return types.Throw(titleErrors, contentErrors)
}

func (r *UpdatePubRequest) Map(authorId uint64) (models.Publication, error) {
	return models.NewPub(r.Title, r.Content, authorId)
}
