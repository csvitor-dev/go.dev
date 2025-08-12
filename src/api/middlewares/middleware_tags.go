package middlewares

import (
	"fmt"
	"slices"
)

type MiddlewareTags struct {
	tags []string
}

func NewTags(tags ...string) *MiddlewareTags {
	for _, tag := range tags {
		if _, exists := allMiddlewares[tag]; !exists {
			panic(fmt.Sprintf("middleware: specified tag '%s' not found", tag))
		}
	}
	return &MiddlewareTags{
		tags: tags,
	}
}

func (m *MiddlewareTags) HasTag(tag string) bool {
	return slices.Contains(m.tags, tag)
}

func (m *MiddlewareTags) AddTag(tag string, isPriority bool) {
	if _, exists := allMiddlewares[tag]; !exists {
		panic(fmt.Sprintf("middleware: specified tag '%s' not found", tag))
	}
	var result []string

	if isPriority {
		result = slices.Insert(m.tags, 0, tag)
	} else {
		result = append(m.tags, tag)
	}
	m.tags = result
}

func (m *MiddlewareTags) AllTags() []string {
	output := m.tags
	slices.Reverse(output)
	return output
}
