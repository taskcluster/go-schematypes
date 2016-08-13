package schematypes

import (
	"math"
	"reflect"
	"regexp"
)

// The Integer struct represents a JSON schema for an integer.
type Integer struct {
	MetaData
	Minimum int64
	Maximum int64
}

// Schema returns a JSON representation of the schema.
func (i Integer) Schema() map[string]interface{} {
	m := i.schema()
	m["type"] = "integer"
	if i.Minimum != math.MinInt64 {
		m["minimum"] = i.Minimum
	}
	if i.Maximum != math.MaxInt64 {
		m["maximum"] = i.Maximum
	}
	return m
}

// Validate the given data, this will return nil if data satisfies this schema.
// Otherwise, Validate(data) returns a list of ValidationIssues.
func (i Integer) Validate(data interface{}) *ValidationError {
	value, ok := data.(float64)
	if !ok || float64(int64(value)) != value {
		return singleIssue("", "Expected an integer at {path}")
	}
	if int64(value) < i.Minimum {
		return singleIssue("", "Integer %d at {path} is less than minimum %d",
			value, i.Minimum,
		)
	}
	if int64(value) > i.Maximum {
		return singleIssue("", "Integer %d at {path} is larger than maximum %d",
			value, i.Maximum,
		)
	}

	return nil
}

// Map takes data, validates and maps it into the target reference.
func (i Integer) Map(data interface{}, target interface{}) error {
	if err := i.Validate(data); err != nil {
		return err
	}

	ptr := reflect.ValueOf(target)
	if ptr.Kind() != reflect.Ptr {
		return ErrTypeMismatch
	}
	val := ptr.Elem()

	switch val.Kind() {
	case reflect.Int8:
		if i.Minimum < math.MinInt8 && i.Maximum > math.MaxInt8 {
			return ErrTypeMismatch
		}
		fallthrough
	case reflect.Int16:
		if i.Minimum < math.MinInt16 && i.Maximum > math.MaxInt16 {
			return ErrTypeMismatch
		}
		fallthrough
	case reflect.Int32, reflect.Int:
		if i.Minimum < math.MinInt32 && i.Maximum > math.MaxInt32 {
			return ErrTypeMismatch
		}
		fallthrough
	case reflect.Int64:
		val.SetInt(int64(data.(float64)))
		return nil
	case reflect.Uint8:
		if i.Minimum < 0 && i.Maximum > math.MaxUint8 {
			return ErrTypeMismatch
		}
		fallthrough
	case reflect.Uint16:
		if i.Minimum < 0 && i.Maximum > math.MaxUint16 {
			return ErrTypeMismatch
		}
		fallthrough
	case reflect.Uint32, reflect.Uint:
		if i.Minimum < 0 && i.Maximum > math.MaxUint32 {
			return ErrTypeMismatch
		}
		fallthrough
	case reflect.Uint64:
		if i.Minimum < 0 {
			return ErrTypeMismatch
		}
		val.SetUint(uint64(data.(float64)))
		return nil
	default:
		return ErrTypeMismatch
	}
}

// Number schema type.
type Number struct {
	MetaData
	Minimum float64
	Maximum float64
}

// Schema returns a JSON representation of the schema.
func (n Number) Schema() map[string]interface{} {
	m := n.schema()
	m["type"] = "number"
	if n.Minimum != -math.MaxFloat64 {
		m["minimum"] = n.Minimum
	}
	if n.Maximum != math.MaxFloat64 {
		m["maximum"] = n.Maximum
	}
	return m
}

// Validate the given data, this will return nil if data satisfies this schema.
// Otherwise, Validate(data) returns a list of ValidationIssues.
func (n Number) Validate(data interface{}) *ValidationError {
	value, ok := data.(float64)
	if !ok {
		return singleIssue("", "Expected a number at {path}")
	}
	if value < n.Minimum {
		return singleIssue("", "Number %d at {path} is less than minimum %d",
			value, n.Minimum,
		)
	}
	if value > n.Maximum {
		return singleIssue("", "Number %d at {path} is larger than maximum %d",
			value, n.Maximum,
		)
	}
	return nil
}

// Map takes data, validates and maps it into the target reference.
func (n Number) Map(data interface{}, target interface{}) error {
	if err := n.Validate(data); err != nil {
		return err
	}

	ptr := reflect.ValueOf(target)
	if ptr.Kind() != reflect.Ptr {
		return ErrTypeMismatch
	}
	val := ptr.Elem()

	switch val.Kind() {
	case reflect.Float32, reflect.Float64:
		val.SetFloat(data.(float64))
		return nil
	default:
		return ErrTypeMismatch
	}
}

// Boolean schema type.
type Boolean struct{ MetaData }

// Schema returns a JSON representation of the schema.
func (b Boolean) Schema() map[string]interface{} {
	m := b.schema()
	m["type"] = "boolean"
	return m
}

// Validate the given data, this will return nil if data satisfies this schema.
// Otherwise, Validate(data) returns a list of ValidationIssues.
func (b Boolean) Validate(data interface{}) *ValidationError {
	if _, ok := data.(bool); !ok {
		return singleIssue("", "Expected a boolean at {path}")
	}
	return nil
}

// Map takes data, validates and maps it into the target reference.
func (b Boolean) Map(data interface{}, target interface{}) error {
	if err := b.Validate(data); err != nil {
		return err
	}

	ptr := reflect.ValueOf(target)
	if ptr.Kind() != reflect.Ptr {
		return ErrTypeMismatch
	}
	val := ptr.Elem()

	switch val.Kind() {
	case reflect.Bool:
		val.SetBool(data.(bool))
		return nil
	default:
		return ErrTypeMismatch
	}
}

// String schema type.
type String struct {
	MetaData
	MinimumLength int
	MaximumLength int
	Pattern       string
}

// Schema returns a JSON representation of the schema.
func (s String) Schema() map[string]interface{} {
	m := s.schema()
	m["type"] = "string"
	if s.MinimumLength != 0 {
		m["minLength"] = s.MinimumLength
	}
	if s.MaximumLength != 0 {
		m["maxLength"] = s.MaximumLength
	}
	if s.Pattern != "" {
		m["pattern"] = s.Pattern
	}
	return m
}

// Validate the given data, this will return nil if data satisfies this schema.
// Otherwise, Validate(data) returns a list of ValidationIssues.
func (s String) Validate(data interface{}) *ValidationError {
	value, ok := data.(string)
	if !ok {
		return singleIssue("", "Expected a string at {path}")
	}

	e := &ValidationError{}

	if s.MinimumLength != 0 && len(value) < s.MinimumLength {
		e.addIssue("",
			"String '%s' at {path} is shorter than minimum %d length allowed",
			value, s.MinimumLength)
	}
	if s.MaximumLength != 0 && len(value) > s.MaximumLength {
		e.addIssue("",
			"String '%s' at {path} is longer than maximum %d length allowed",
			value, s.MaximumLength)
	}
	if s.Pattern != "" {
		if match, _ := regexp.MatchString(s.Pattern, value); !match {
			e.addIssue("", "String '%s' doesn't match regular expression '%s'",
				value, s.Pattern)
		}
	}

	if len(e.issues) > 0 {
		return e
	}
	return nil
}

// Map takes data, validates and maps it into the target reference.
func (s String) Map(data interface{}, target interface{}) error {
	if err := s.Validate(data); err != nil {
		return err
	}

	ptr := reflect.ValueOf(target)
	if ptr.Kind() != reflect.Ptr {
		return ErrTypeMismatch
	}
	val := ptr.Elem()

	switch val.Kind() {
	case reflect.String:
		val.SetString(data.(string))
		return nil
	default:
		return ErrTypeMismatch
	}
}

// StringEnum schema type for enums of strings.
type StringEnum struct {
	MetaData
	Values []string
}

// Schema returns a JSON representation of the schema.
func (s StringEnum) Schema() map[string]interface{} {
	m := s.schema()
	m["type"] = "string"
	m["enum"] = s.Values
	return m
}

// Validate the given data, this will return nil if data satisfies this schema.
// Otherwise, Validate(data) returns a list of ValidationIssues.
func (s StringEnum) Validate(data interface{}) *ValidationError {
	value, ok := data.(string)
	if !ok {
		return singleIssue("", "Expected a string at {path}")
	}

	if !stringContains(s.Values, value) {
		e := &ValidationError{}
		e.addIssue("",
			"Value '%s' at {path} is not valid for the enum with options: %v",
			value, s.Values)
		return e
	}

	return nil
}

// Map takes data, validates and maps it into the target reference.
func (s StringEnum) Map(data interface{}, target interface{}) error {
	if err := s.Validate(data); err != nil {
		return err
	}

	ptr := reflect.ValueOf(target)
	if ptr.Kind() != reflect.Ptr {
		return ErrTypeMismatch
	}
	val := ptr.Elem()

	switch val.Kind() {
	case reflect.String:
		val.SetString(data.(string))
		return nil
	default:
		return ErrTypeMismatch
	}
}
