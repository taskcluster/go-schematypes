package schematypes

import "testing"

func TestSchema(t *testing.T) {
	var aInterface interface{}
	schema := `{
    "type": "object",
    "title": "my-title-3",
    "description": "my-description-3",
    "properties": {
      "obj": {
        "type": "object",
        "title": "my-title-1",
        "description": "my-description-1",
        "properties": {
          "int": {
            "type": "integer",
            "title": "my-title-2",
            "description": "my-description-2",
            "minimum": -240,
            "maximum": 240
          }
        },
        "additionalProperties": false,
        "required": ["int"]
      }
    },
    "additionalProperties": false,
    "required": ["obj"]
  }`
	s, err := NewSchema(schema)
	if err != nil {
		panic(err)
	}
	testCase{
		Schema: s,
		Match:  schema,
		Valid: []string{
			"{\"obj\": {\"int\": 4}}",
		},
		Invalid: []string{
			"[]", "{\"int\": 4}",
		},
		TypeMatch: []interface{}{
			&aInterface,
		},
		TypeMismatch: []interface{}{
			&struct {
				Int int `json:"int"`
			}{},
			&struct {
				Obj **struct {
					Int int `json:"int"`
				} `json:"obj"`
			}{},
			&struct {
				Obj *struct {
					Int int `json:"wrong"`
				} `json:"obj"`
			}{},
			&struct {
				Obj *struct {
					Int int `json:"int"`
				} `json:"wrong"`
			}{},
		},
	}.Test(t)
}
