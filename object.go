package schematypes

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
)

// Properties defines the properties for a object schema.
type Properties map[string]Schema

// Object specifies schema for an object.
type Object struct {
	MetaData
	Properties           Properties
	AdditionalProperties bool
	Required             []string
}

// Schema returns a JSON representation of the schema.
func (o Object) Schema() map[string]interface{} {
	m := o.schema()
	m["type"] = "object"
	if len(o.Properties) > 0 {
		props := make(map[string]map[string]interface{})
		for prop, schema := range o.Properties {
			props[prop] = schema.Schema()
		}
		m["properties"] = props
	}
	if !o.AdditionalProperties {
		m["additionalProperties"] = o.AdditionalProperties
	}
	if len(o.Required) > 0 {
		m["required"] = o.Required
	}
	return m
}

var identifierPattern = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")

func formatKeyPath(key string) string {
	if identifierPattern.MatchString(key) {
		return "." + key
	}
	j, _ := json.Marshal(key)
	return "[" + string(j) + "]"
}

// Validate the given data, this will return nil if data satisfies this schema.
// Otherwise, Validate(data) returns a list of ValidationIssues.
func (o Object) Validate(data interface{}) *ValidationError {
	value, ok := data.(map[string]interface{})
	if !ok {
		return singleIssue("", "Expected object type at {path}")
	}

	e := ValidationError{}

	// Test properties
	for p, s := range o.Properties {
		if err := s.Validate(value[p]); err != nil {
			e.addIssuesWithPrefix(err, formatKeyPath(p))
		}
	}

	// Test for additional properties
	if !o.AdditionalProperties {
		for key := range value {
			if _, ok := o.Properties[key]; !ok {
				e.addIssue(formatKeyPath(key), "Additional property '%s' not allowed at {path}", key)
			}
		}
	}

	// Test required properties
	for _, key := range o.Required {
		if _, ok := value[key]; !ok {
			e.addIssue(formatKeyPath(key), "Required property '%s' is missing at {path}", key)
		}
	}

	if len(e.issues) > 0 {
		return &e
	}
	return nil
}

// Map takes data, validates and maps it into the target reference.
func (o Object) Map(data, target interface{}) error {
	if err := o.Validate(data); err != nil {
		return err
	}

	// Ensure that we have a pointer as input
	ptr := reflect.ValueOf(target)
	if ptr.Kind() != reflect.Ptr {
		return ErrTypeMismatch
	}
	val := ptr.Elem()

	// Use mapStruct if we have a struct type
	if val.Kind() == reflect.Struct {
		return o.mapStruct(data.(map[string]interface{}), val)
	}

	// TODO: Add support for mapping to a map[string]interface{}
	// TODO: Add support for mapping to a map[string]<T>

	return ErrTypeMismatch
}

func jsonTag(field reflect.StructField) string {
	j := field.Tag.Get("json")
	if strings.HasSuffix(j, ",omitempty") {
		return j[:len(j)-10]
	}
	return j
}

func hasStructTag(t reflect.Type, tag string) bool {
	N := t.NumField()
	for i := 0; i < N; i++ {
		f := t.Field(i)
		j := jsonTag(f)
		if j == tag {
			return true
		} else if j == "" && f.Name == tag {
			return true
		}
	}
	return false
}

var typeOfEmptyInterface = reflect.TypeOf((*interface{})(nil)).Elem()

func (o Object) mapStruct(data map[string]interface{}, target reflect.Value) error {
	t := target.Type()

	// We have a type mismatch if there isn't fields for the values declared
	for key := range o.Properties {
		if !hasStructTag(t, key) {
			return ErrTypeMismatch
		}
	}

	for _, key := range o.Required {
		if !hasStructTag(t, key) {
			return ErrTypeMismatch
		}
	}

	N := t.NumField()
	for i := 0; i < N; i++ {
		// Find field and json tag
		f := t.Field(i)
		tag := jsonTag(f)

		// Find value, if there is one
		value, ok := data[tag]
		if !ok {
			continue // We've already validated, so no need to check required
		}

		// Find schema for property, ignore if there is none
		s := o.Properties[tag]
		if s == nil {
			continue
		}

		var targetValue reflect.Value
		if f.Type.Kind() == reflect.Ptr {
			targetValue = reflect.New(f.Type.Elem())
			target.Field(i).Set(targetValue)
		} else if f.Type == typeOfEmptyInterface {
			target.Field(i).Set(reflect.ValueOf(value))
			continue
		} else {
			targetValue = target.Field(i).Addr()
		}

		// Map value to field
		err := s.Map(value, targetValue.Interface())
		if err != nil {
			return err
		}
	}

	return nil
}
