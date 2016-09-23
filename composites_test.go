package schematypes

import "testing"

var any interface{}

func TestAnyOf(t *testing.T) {
	testCase{
		Schema: AnyOf{
			Integer{
				Minimum: -240,
				Maximum: 240,
			},
			String{},
			Number{
				Minimum: -240,
				Maximum: 240,
			},
		},
		Match: `{
      "anyOf": [
        {
          "type": "integer",
          "minimum": -240,
          "maximum": 240
        }, {
          "type": "string"
        }, {
          "type": "number",
          "minimum": -240,
          "maximum": 240
        }
      ]
    }`,
		Valid: []string{
			"2", "\"hello-world\"",
		},
		Invalid: []string{
			"[]", "{}", "300",
		},
		TypeMatch: []interface{}{
			&any,
		},
		TypeMismatch: []interface{}{
			&struct {
				Int int `json:"wrong"`
			}{},
			pInt32,
			&[]int8{},
			&[]int32{},
			&[]int{},
			&[]struct{}{},
			&[]string{},
		},
	}.Test(t)
}

func TestOneOf(t *testing.T) {
	testCase{
		Schema: OneOf{
			Integer{
				Minimum: -240,
				Maximum: 240,
			},
			String{},
			Number{
				Minimum: 0,
				Maximum: 240,
			},
		},
		Match: `{
      "oneOf": [
        {
          "type": "integer",
          "minimum": -240,
          "maximum": 240
        }, {
          "type": "string"
        }, {
          "type": "number",
          "minimum": 0,
          "maximum": 240
        }
      ]
    }`,
		Valid: []string{
			"-2", "\"hello-world\"",
		},
		Invalid: []string{
			"[]", "{}", "2", "300",
		},
		TypeMatch: []interface{}{
			&any,
		},
		TypeMismatch: []interface{}{
			&struct {
				Int int `json:"wrong"`
			}{},
			pInt32,
			&[]int8{},
			&[]int32{},
			&[]int{},
			&[]struct{}{},
			&[]string{},
		},
	}.Test(t)
}

func TestAllOf(t *testing.T) {
	testCase{
		Schema: AllOf{
			Integer{
				Minimum: -240,
				Maximum: 240,
			},
			Number{
				Minimum: 0,
				Maximum: 240,
			},
		},
		Match: `{
      "allOf": [
        {
          "type": "integer",
          "minimum": -240,
          "maximum": 240
        }, {
          "type": "number",
          "minimum": 0,
          "maximum": 240
        }
      ]
    }`,
		Valid: []string{
			"2", "100",
		},
		Invalid: []string{
			"[]", "{}", "300", "-2",
		},
		TypeMatch: []interface{}{
			&any,
		},
		TypeMismatch: []interface{}{
			&struct {
				Int int `json:"wrong"`
			}{},
			pInt32,
			&[]int8{},
			&[]int32{},
			&[]int{},
			&[]struct{}{},
			&[]string{},
		},
	}.Test(t)
}
