package publication

import (
	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type CreatePubRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *CreatePubRequest) Validate() types.RequestValidationGuard {
	titleErrors := validations.NewString(r.Title, "title").IsNotEmpty().MaxLength(100).Result()

	contentErrors := validations.NewString(r.Content, "content").IsNotEmpty().MaxLength(300).Result()

	return types.Throw(titleErrors, contentErrors)
}

func (r *CreatePubRequest) Map(authorId uint64) (models.Publication, error) {
	return models.NewPub(r.Title, r.Content, authorId)
}
