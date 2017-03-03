package schematypes

import "fmt"

// stringContains returns true if list contains element
func stringContains(list []string, element string) bool {
	for _, s := range list {
		if s == element {
			return true
		}
	}
	return false
}

// MustValidate panics if data doesn't validate against schema
func MustValidate(schema Schema, data interface{}) {
	err := schema.Validate(data)
	if err != nil {
		panic(fmt.Sprintf(
			"ValidationError schema: %#v doesn't match data: %#v\n%s",
			schema, data, err,
		))
	}
}

// MustMap will map data into target using schema and panic, if it returns
// ErrTypeMismatch
func MustMap(schema Schema, data, target interface{}) error {
	err := schema.Map(data, target)
	if err == ErrTypeMismatch {
		panic(fmt.Sprintf(
			"ErrTypeMismatch, target type: %#v doesn't match schema: %#v",
			target, schema,
		))
	}
	return err
}

// MustValidateAndMap panics if data doesn't validate or maps into target
func MustValidateAndMap(schema Schema, data, target interface{}) {
	err := schema.Map(data, target)
	if err == ErrTypeMismatch {
		panic(fmt.Sprintf(
			"ErrTypeMismatch, target type: %#v doesn't match schema: %#v",
			target, schema,
		))
	}
	if err != nil {
		panic(fmt.Sprintf(
			"ValidationError schema: %#v doesn't match data: %#v\n%s",
			schema, data, err,
		))
	}
}
