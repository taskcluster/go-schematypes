package schematypes

import (
	"errors"
	"fmt"
	"strings"
)

// ErrTypeMismatch is returned when trying to map to a type that doesn't match
// or which isn't writable (for example passed by value and not pointer).
var ErrTypeMismatch = errors.New("Type does not match the schema")

// A ValidationIssue is any error found validating a JSON object.
type ValidationIssue struct {
	message string
	path    string
}

// String returns a human readable string representation of the issue
func (v *ValidationIssue) String() string {
	return strings.Replace(v.message, "{path}", v.path, -1)
}

// Path returns the a path to the issue, on the form:
//   rootName.dictionary["other-key"].array[44].property
func (v *ValidationIssue) Path() string {
	return v.path
}

// prefix will add a prefix to the path to property that had an issue.
func (v *ValidationIssue) prefix(prefix string, args ...interface{}) ValidationIssue {
	return ValidationIssue{
		message: v.message,
		path:    fmt.Sprintf(prefix, args...) + v.path,
	}
}

// ValidationError represents a validation failure as a list of validation
// issues.
type ValidationError struct {
	issues []ValidationIssue
}

// Issues returns the validation issues with given rootName as the start of the
// Path or "root" if rootName is the empty string.
func (e *ValidationError) Issues(rootName string) []ValidationIssue {
	if rootName == "" {
		rootName = "root"
	}
	issues := make([]ValidationIssue, len(e.issues))
	for i, issue := range e.issues {
		issues[i] = issue.prefix("%s", rootName)
	}
	return issues
}

func (e *ValidationError) addIssue(path, message string, args ...interface{}) {
	e.issues = append(e.issues, ValidationIssue{
		message: fmt.Sprintf(message, args...),
		path:    path,
	})
}

func (e *ValidationError) addIssues(err *ValidationError) {
	e.issues = append(e.issues, err.issues...)
}

func (e *ValidationError) addIssuesWithPrefix(err error, prefix string, args ...interface{}) {
	if err == nil {
		return
	}
	if err, ok := err.(*ValidationError); ok {
		for _, issue := range err.issues {
			e.issues = append(e.issues, issue.prefix(prefix, args...))
		}
	} else {
		issue := ValidationIssue{
			message: fmt.Sprintf("Error: %s at {path}", err.Error()),
			path:    "",
		}
		e.issues = append(e.issues, issue.prefix(prefix, args...))
	}
}

func singleIssue(path, message string, args ...interface{}) *ValidationError {
	e := &ValidationError{}
	e.addIssue(path, message, args...)
	return e
}

func (e *ValidationError) Error() string {
	msg := "ValidationError: "
	for _, issue := range e.issues {
		msg += strings.Replace(issue.message, "{path}", "root"+issue.path, -1) + ", "
	}
	return msg
}
