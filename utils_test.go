package schematypes

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func fmtPanic(a ...interface{}) {
	panic(fmt.Sprintln(a...))
}

func nilOrPanic(err error, a ...interface{}) {
	if err != nil {
		fmtPanic(append(a, err)...)
	}
}

func evalNilOrPanic(f func() error, a ...interface{}) {
	nilOrPanic(f(), a...)
}

func assert(condition bool, a ...interface{}) {
	if !condition {
		fmtPanic(a...)
	}
}

func assertJSON(val interface{}, jsonString string, a ...interface{}) {
	// Marshal and unmarshall to get int to float64 form
	data, err := json.Marshal(val)
	nilOrPanic(err, "Input value can't be serialized to JSON, err:", err)
	var v1 interface{}
	err = json.Unmarshal(data, &v1)
	nilOrPanic(err, "Internal error!")

	var v2 interface{}
	err = json.Unmarshal([]byte(jsonString), &v2)
	nilOrPanic(err, "Internal test error, jsonString isn't valid json, err: ", err)

	assert(reflect.DeepEqual(v1, v2), a...)
}

func testValidate(t *testing.T, s Schema, jsonString string, valid bool) {
	var val interface{}
	err := json.Unmarshal([]byte(jsonString), &val)
	nilOrPanic(err, "Internal test error, jsonString isn't valid json, err: ", err)

	fail := false
	if valid {
		e := s.Validate(val)
		if e != nil {
			fmt.Println("Unexpected validation error: ", e)
			fail = true
		}
	} else {
		if s.Validate(val) == nil {
			fmt.Println("Expected some validation error")
			fail = true
		}
	}

	if fail {
		fmt.Println("                  <-- In test case: ")
		fmt.Println("Expected valid: ", valid)
		fmt.Println("Data:   ", jsonString)
		fmt.Println("Schema: ", s.Schema())
		fmt.Println("---------------")
		t.Fail()
	}
}

func testMap(t *testing.T, s Schema, jsonString string, target interface{}, valid, typeMatch bool) {
	var v interface{}
	err := json.Unmarshal([]byte(jsonString), &v)
	nilOrPanic(err, "Internal test error, jsonString isn't valid json, err: ", err)

	err = s.Map(v, target)
	fail := false
	if !valid {
		if err == nil {
			fmt.Println("Expected a validation error, got no error")
			fail = true
		}
		if !typeMatch && err == ErrTypeMismatch {
			fmt.Println("Expected validation error got ErrTypeMismatch")
			fail = true
		}
	} else {
		if !typeMatch {
			if err != ErrTypeMismatch {
				fmt.Println("Expected ErrTypeMismatch, but got err = ", err)
				fail = true
			}
		} else {
			if err != nil {
				fmt.Println("Unexpected error: ", err)
				fail = true
			}
		}
	}

	if fail {
		fmt.Println("                  <-- In test case: ")
		fmt.Println("Expected valid:      ", valid)
		fmt.Println("Expected type match: ", typeMatch)
		fmt.Println("Data:   ", jsonString)
		fmt.Println("Target: ", reflect.TypeOf(target))
		fmt.Println("Schema: ", s.Schema())
		fmt.Println("---------------")
		t.Fail()
	}
}

// Just testing our test utility
func TestAssertJSON(t *testing.T) {
	assertJSON(map[string]interface{}{
		"integer": 21,
		"list":    []string{"a", "b"},
		"obj": map[string]interface{}{
			"true":  true,
			"false": false,
		},
	}, `{
		"integer": 21,
		"list": ["a", "b"],
		"obj": {
			"true": true,
			"false": false
		}
	}`, "Expected JSON string to match object")
}

type testCase struct {
	Schema       Schema
	Match        string
	Valid        []string
	Invalid      []string
	TypeMatch    []interface{}
	TypeMismatch []interface{}
}

func (c testCase) Test(t *testing.T) {
	assertJSON(c.Schema.Schema(), c.Match, "Schema() wasn't equal to Match")

	for _, s := range c.Valid {
		testValidate(t, c.Schema, s, true)
	}
	for _, s := range c.Invalid {
		testValidate(t, c.Schema, s, false)
	}

	for _, target := range c.TypeMatch {
		for _, s := range c.Valid {
			testMap(t, c.Schema, s, target, true, true)
		}
		for _, s := range c.Invalid {
			testMap(t, c.Schema, s, target, false, true)
		}
	}

	for _, target := range c.TypeMismatch {
		for _, s := range c.Valid {
			testMap(t, c.Schema, s, target, true, false)
		}
	}
}
