package restaurant

import "fmt"

// Tag .
type Tag string

// NewTag .
func NewTag(tag string) (Tag, error) {
	if tag == "" {
		return "", fmt.Errorf("invalid tag")
	}
	return Tag(tag), nil
}

// String .
func (tag Tag) String() string {
	return string(tag)
}

// Tags .
type Tags []Tag

// NewTags .
func NewTags(ts []string) (Tags, error) {
	tags := make(Tags, len(ts))
	for i, t := range ts {
		tag, err := NewTag(t)
		if err != nil {
			return nil, err
		}
		tags[i] = tag
	}
	return tags, nil
}

// Strings .
func (ts Tags) Strings() []string {
	tags := make([]string, len(ts))
	for i, t := range ts {
		tags[i] = t.String()
	}
	return tags
}
