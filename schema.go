package schematypes

// A Schema is implemented by any object that can represent a JSON schema.
type Schema interface {
	Schema() map[string]interface{}
	Validate(data interface{}) *ValidationError
	Map(data, target interface{}) error
}

/*
func Test() {
	schema := Object{
		Properties: Properties{
			"integer": Integer{},
		},
		AdditionalProperties: false,
		Required: []string{
			"--",
			"dd",
		},
	}
	err := schema.Validate(map[string]interface{}{
		"integer": 42,
	})
	type MyType struct {
		Integer int `json:"integer,omitempty"`
	}
	var instance MyType
	err := schema.Map(data, &instance)
	if err == schema.ErrTypeMismatch {
		panic("...")
	}
	if err != nil {
		return NewMalformedPayloadError(err.String())
	}
}
*/

// An AnyOf instance represents the anyOf JSON schema construction.
type AnyOf []Schema

// A OneOf instance represents the oneOf JSON schema construction.
type OneOf []Schema

// An AllOf instance represents the allOf JSON schema construction.
type AllOf []Schema
