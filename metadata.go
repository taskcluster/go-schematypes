package schematypes

// MetaData contains meta-data properties common to all schema types.
type MetaData struct {
	Title       string
	Description string
}

func (s *MetaData) schema() map[string]interface{} {
	m := make(map[string]interface{})
	if s.Title != "" {
		m["title"] = s.Title
	}
	if s.Description != "" {
		m["description"] = s.Description
	}
	return m
}

func makeMetaData(title, description string) map[string]interface{} {
	m := make(map[string]interface{})
	if title != "" {
		m["title"] = title
	}
	if description != "" {
		m["description"] = description
	}
	return m
}
