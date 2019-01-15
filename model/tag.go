// Package model contains all the entities
package model

// The tag entity
type Tag struct {
	// tag name
	Tag string `json:"tag"`

	// tag count across articles
	Count uint64 `json:"count"`

	// list of articles
	Articles []uint64 `json:"articles"`

	// list of cooccuring tags
	RelatedTags []string `json:"related_tags"`
}
