package schematypes

import "testing"

func TestIntegerArray(t *testing.T) {
	testCase{
		Schema: Array{
			MetaData: MetaData{
				Title:       "my-title-1",
				Description: "my-description-1",
			},
			Items: Integer{
				MetaData: MetaData{
					Title:       "my-title-2",
					Description: "my-description-2",
				},
				Minimum: -240,
				Maximum: 240,
			},
		},
		Match: `{
      "type": "array",
      "title": "my-title-1",
      "description": "my-description-1",
      "items": {
        "type": "integer",
        "title": "my-title-2",
        "description": "my-description-2",
        "minimum": -240,
        "maximum": 240
      }
    }`,
		Valid: []string{
			"[4,67,76]",
			"[4,67,76,67]",
		},
		Invalid: []string{
			"{}", "56", "34.3", "\"dfsdf\"",
			"[4,67,7.66]", "[4,67,500]",
		},
		TypeMatch: []interface{}{
			&[]int32{},
			&[]int{},
		},
		TypeMismatch: []interface{}{
			&[]int8{},
			&[]struct{}{},
			&[]string{},
		},
	}.Test(t)
}

func TestIntegerArrayEmpty(t *testing.T) {
	testCase{
		Schema: Array{
			MetaData: MetaData{
				Title:       "my-title-1",
				Description: "my-description-1",
			},
			Items: Integer{
				MetaData: MetaData{
					Title:       "my-title-2",
					Description: "my-description-2",
				},
				Minimum: -240,
				Maximum: 240,
			},
		},
		Match: `{
      "type": "array",
      "title": "my-title-1",
      "description": "my-description-1",
      "items": {
        "type": "integer",
        "title": "my-title-2",
        "description": "my-description-2",
        "minimum": -240,
        "maximum": 240
      }
    }`,
		Valid: []string{
			"[]",
		},
		Invalid: []string{
			"{}", "56", "34.3", "\"dfsdf\"",
			"[4,67,7.66]", "[4,67,500]",
		},
		TypeMatch: []interface{}{
			&[]int32{},
		},
		TypeMismatch: []interface{}{
			// We don't force ErrTypeMismatch, if the array is empty...
			&struct{}{},
		},
	}.Test(t)
}

func TestUniqueIntegerArray(t *testing.T) {
	testCase{
		Schema: Array{
			MetaData: MetaData{
				Title:       "my-title-1",
				Description: "my-description-1",
			},
			Items: Integer{
				MetaData: MetaData{
					Title:       "my-title-2",
					Description: "my-description-2",
				},
				Minimum: -240,
				Maximum: 240,
			},
			Unique: true,
		},
		Match: `{
      "type": "array",
      "title": "my-title-1",
      "description": "my-description-1",
      "items": {
        "type": "integer",
        "title": "my-title-2",
        "description": "my-description-2",
        "minimum": -240,
        "maximum": 240
      },
      "uniqueItems": true
    }`,
		Valid: []string{
			"[4,67,76]",
		},
		Invalid: []string{
			"{}", "56", "34.3", "\"dfsdf\"",
			"[4,67,7.66]", "[4,67,500]",
			"[4,67,76,67]",
		},
		TypeMatch: []interface{}{
			&[]int32{},
			&[]int{},
		},
		TypeMismatch: []interface{}{
			&[]int8{},
			&[]struct{}{},
			&[]string{},
		},
	}.Test(t)
}
