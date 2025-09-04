package publication

import (
	"github.com/csvitor-dev/go.dev/internal/models"
	"github.com/csvitor-dev/go.dev/types"
	"github.com/csvitor-dev/go.dev/utils/validations"
)

type CreatePubRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *CreatePubRequest) Validate() types.RequestValidationGuard {
	title := validations.NewString(r.Title, "title").IsNotEmpty().MaxLength(100).TrimRefine()
	content := validations.NewString(r.Content, "content").IsNotEmpty().MaxLength(300).TrimRefine()

	return types.Throw(title, content)
}

func (r *CreatePubRequest) Map(authorId uint64) (models.Publication, error) {
	return models.NewPub(r.Title, r.Content, authorId)
}
