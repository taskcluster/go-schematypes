package schematypes

import "testing"

func TestObject(t *testing.T) {
	testCase{
		Schema: Object{
			MetaData: MetaData{
				Title:       "my-title-1",
				Description: "my-description-1",
			},
			Properties: Properties{
				"int": Integer{
					MetaData: MetaData{
						Title:       "my-title-2",
						Description: "my-description-2",
					},
					Minimum: -240,
					Maximum: 240,
				},
			},
			Required: []string{"int"},
		},
		Match: `{
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
    }`,
		Valid: []string{
			"{\"int\": 4}",
		},
		Invalid: []string{
			"[]", "{}",
		},
		TypeMatch: []interface{}{
			&struct {
				Int int `json:"int"`
			}{},
		},
		TypeMismatch: []interface{}{
			&struct {
				Int int `json:"wrong"`
			}{},
			pInt32,
			&[]int8{},
			&[]struct{}{},
			&[]string{},
			&[]int32{},
			&[]int{},
		},
	}.Test(t)
}

func TestOptionalPropertyObject(t *testing.T) {
	testCase{
		Schema: Object{
			MetaData: MetaData{
				Title:       "my-title-1",
				Description: "my-description-1",
			},
			Properties: Properties{
				"int": Integer{
					MetaData: MetaData{
						Title:       "my-title-2",
						Description: "my-description-2",
					},
					Minimum: -240,
					Maximum: 240,
				},
			},
		},
		Match: `{
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
      "additionalProperties": false
    }`,
		Valid: []string{
			"{\"int\": 4}", "{}",
		},
		Invalid: []string{
			"[]",
		},
		TypeMatch: []interface{}{
			&struct {
				Int int `json:"int"`
			}{},
		},
		TypeMismatch: []interface{}{
			&struct {
				Int int `json:"wrong"`
			}{},
			pInt32,
			&[]int8{},
			&[]struct{}{},
			&[]string{},
			&[]int32{},
			&[]int{},
		},
	}.Test(t)
}

func TestNestedObject(t *testing.T) {
	testCase{
		Schema: Object{
			MetaData: MetaData{
				Title:       "my-title-3",
				Description: "my-description-3",
			},
			Properties: Properties{
				"obj": Object{
					MetaData: MetaData{
						Title:       "my-title-1",
						Description: "my-description-1",
					},
					Properties: Properties{
						"int": Integer{
							MetaData: MetaData{
								Title:       "my-title-2",
								Description: "my-description-2",
							},
							Minimum: -240,
							Maximum: 240,
						},
					},
					Required: []string{"int"},
				},
			},
			Required: []string{"obj"},
		},
		Match: `{
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
		}`,
		Valid: []string{
			"{\"obj\": {\"int\": 4}}",
		},
		Invalid: []string{
			"[]", "{\"int\": 4}",
		},
		TypeMatch: []interface{}{
			&struct {
				Obj struct {
					Int int `json:"int"`
				} `json:"obj"`
			}{},
			&struct {
				Obj *struct {
					Int int `json:"int"`
				} `json:"obj"`
			}{},
			&struct {
				Obj interface{} `json:"obj"`
			}{},
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
