package publication

import (
	"github.com/csvitor-dev/go.dev/internal/models"
	"github.com/csvitor-dev/go.dev/types"
	"github.com/csvitor-dev/go.dev/utils/validations"
)

type UpdatePubRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

func (r *UpdatePubRequest) Validate() types.RequestValidationGuard {
	title := validations.NewString(r.Title, "title").IsOptional().MaxLength(100)
	content := validations.NewString(r.Content, "content").IsOptional().MaxLength(300)

	if optional := validations.
		AllOptionalExpressionsAreValid(
			title,
			content,
		); optional != nil {
		return types.Throw(optional)
	}
	return types.Throw(title, content)
}

func (r *UpdatePubRequest) Map(authorId uint64) (models.Publication, error) {
	return models.NewPub(r.Title, r.Content, authorId)
}
