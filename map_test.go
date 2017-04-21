package schematypes

import "testing"

func TestMapInteger(t *testing.T) {
	var smap map[string]int
	testCase{
		Schema: Map{
			MetaData: MetaData{
				Title:       "my-title-1",
				Description: "my-description-1",
			},
			Values: Integer{
				Title:       "my-title-2",
				Description: "my-description-2",
				Minimum:     -240,
				Maximum:     240,
			},
		},
		Match: `{
      "type": "object",
      "title": "my-title-1",
      "description": "my-description-1",
      "additionalProperties": {
        "type": "integer",
        "title": "my-title-2",
        "description": "my-description-2",
        "minimum": -240,
        "maximum": 240
      }
    }`,
		Valid: []string{
			"{}",
			"{\"test\":4, \"test2\":67, \"test3\":76}",
			"{\"test\":4, \"test2\":67, \"test3\":76, \"hello world\": 67}",
		},
		Invalid: []string{
			"56", "34.3", "\"dfsdf\"",
			"[4,67,7.66]", "[4,67,500]",
			"{\"test\":4, \"test2\":67, \"test3\":500}",
		},
		TypeMatch: []interface{}{
			&smap,
		},
		TypeMismatch: []interface{}{
			&[]int32{},
			&[]int{},
			&[]int8{},
			&[]struct{}{},
			&[]string{},
		},
	}.Test(t)
}

func TestMapString(t *testing.T) {
	var smap map[string]string
	testCase{
		Schema: Map{
			MetaData: MetaData{
				Title:       "my-title-1",
				Description: "my-description-1",
			},
			Values: String{
				Title:       "my-title-2",
				Description: "my-description-2",
			},
		},
		Match: `{
      "type": "object",
      "title": "my-title-1",
      "description": "my-description-1",
      "additionalProperties": {
        "type": "string",
        "title": "my-title-2",
        "description": "my-description-2"
      }
    }`,
		Valid: []string{
			"{}",
			"{\"test\":\"4 \", \"test2 lala \":\"67\", \"test3\":\"hello world\"}",
		},
		Invalid: []string{
			"56", "34.3", "\"dfsdf\"",
			"[4,67,7.66]", "[4,67,500]",
			"{\"test\":4, \"test2\":67, \"test3\":500}",
		},
		TypeMatch: []interface{}{
			&smap,
		},
		TypeMismatch: []interface{}{
			&[]int32{},
			&[]int{},
			&[]int8{},
			&[]struct{}{},
			&[]string{},
		},
	}.Test(t)
}
