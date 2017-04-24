package schematypes

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
