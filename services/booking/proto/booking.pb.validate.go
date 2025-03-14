// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services/booking/proto/booking.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaskRequestMultiError, or nil if none found.
func (m *CreateTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Detail

	if len(errors) > 0 {
		return CreateTaskRequestMultiError(errors)
	}

	return nil
}

// CreateTaskRequestMultiError is an error wrapping multiple validation errors
// returned by CreateTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaskRequestMultiError) AllErrors() []error { return m }

// CreateTaskRequestValidationError is the validation error returned by
// CreateTaskRequest.Validate if the designated constraints aren't met.
type CreateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskRequestValidationError) ErrorName() string {
	return "CreateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskRequestValidationError{}

// Validate checks the field values on CreateTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaskResponseMultiError, or nil if none found.
func (m *CreateTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaskResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTaskResponseMultiError(errors)
	}

	return nil
}

// CreateTaskResponseMultiError is an error wrapping multiple validation errors
// returned by CreateTaskResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaskResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaskResponseMultiError) AllErrors() []error { return m }

// CreateTaskResponseValidationError is the validation error returned by
// CreateTaskResponse.Validate if the designated constraints aren't met.
type CreateTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskResponseValidationError) ErrorName() string {
	return "CreateTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskResponseValidationError{}

// Validate checks the field values on AcceptTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AcceptTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AcceptTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AcceptTaskRequestMultiError, or nil if none found.
func (m *AcceptTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AcceptTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TaskId

	if len(errors) > 0 {
		return AcceptTaskRequestMultiError(errors)
	}

	return nil
}

// AcceptTaskRequestMultiError is an error wrapping multiple validation errors
// returned by AcceptTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type AcceptTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AcceptTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AcceptTaskRequestMultiError) AllErrors() []error { return m }

// AcceptTaskRequestValidationError is the validation error returned by
// AcceptTaskRequest.Validate if the designated constraints aren't met.
type AcceptTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AcceptTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AcceptTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AcceptTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AcceptTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AcceptTaskRequestValidationError) ErrorName() string {
	return "AcceptTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AcceptTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAcceptTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AcceptTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AcceptTaskRequestValidationError{}

// Validate checks the field values on AcceptTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AcceptTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AcceptTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AcceptTaskResponseMultiError, or nil if none found.
func (m *AcceptTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AcceptTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsSuccess

	if len(errors) > 0 {
		return AcceptTaskResponseMultiError(errors)
	}

	return nil
}

// AcceptTaskResponseMultiError is an error wrapping multiple validation errors
// returned by AcceptTaskResponse.ValidateAll() if the designated constraints
// aren't met.
type AcceptTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AcceptTaskResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AcceptTaskResponseMultiError) AllErrors() []error { return m }

// AcceptTaskResponseValidationError is the validation error returned by
// AcceptTaskResponse.Validate if the designated constraints aren't met.
type AcceptTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AcceptTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AcceptTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AcceptTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AcceptTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AcceptTaskResponseValidationError) ErrorName() string {
	return "AcceptTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AcceptTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAcceptTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AcceptTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AcceptTaskResponseValidationError{}

// Validate checks the field values on ConfirmTaskerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfirmTaskerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfirmTaskerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfirmTaskerRequestMultiError, or nil if none found.
func (m *ConfirmTaskerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfirmTaskerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TaskId

	// no validation rules for TaskerId

	if len(errors) > 0 {
		return ConfirmTaskerRequestMultiError(errors)
	}

	return nil
}

// ConfirmTaskerRequestMultiError is an error wrapping multiple validation
// errors returned by ConfirmTaskerRequest.ValidateAll() if the designated
// constraints aren't met.
type ConfirmTaskerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfirmTaskerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfirmTaskerRequestMultiError) AllErrors() []error { return m }

// ConfirmTaskerRequestValidationError is the validation error returned by
// ConfirmTaskerRequest.Validate if the designated constraints aren't met.
type ConfirmTaskerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfirmTaskerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfirmTaskerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfirmTaskerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfirmTaskerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfirmTaskerRequestValidationError) ErrorName() string {
	return "ConfirmTaskerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ConfirmTaskerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfirmTaskerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfirmTaskerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfirmTaskerRequestValidationError{}

// Validate checks the field values on ConfirmTaskerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfirmTaskerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfirmTaskerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfirmTaskerResponseMultiError, or nil if none found.
func (m *ConfirmTaskerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfirmTaskerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsSuccess

	if len(errors) > 0 {
		return ConfirmTaskerResponseMultiError(errors)
	}

	return nil
}

// ConfirmTaskerResponseMultiError is an error wrapping multiple validation
// errors returned by ConfirmTaskerResponse.ValidateAll() if the designated
// constraints aren't met.
type ConfirmTaskerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfirmTaskerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfirmTaskerResponseMultiError) AllErrors() []error { return m }

// ConfirmTaskerResponseValidationError is the validation error returned by
// ConfirmTaskerResponse.Validate if the designated constraints aren't met.
type ConfirmTaskerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfirmTaskerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfirmTaskerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfirmTaskerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfirmTaskerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfirmTaskerResponseValidationError) ErrorName() string {
	return "ConfirmTaskerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ConfirmTaskerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfirmTaskerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfirmTaskerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfirmTaskerResponseValidationError{}

// Validate checks the field values on GetTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTaskRequestMultiError,
// or nil if none found.
func (m *GetTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TaskId

	if len(errors) > 0 {
		return GetTaskRequestMultiError(errors)
	}

	return nil
}

// GetTaskRequestMultiError is an error wrapping multiple validation errors
// returned by GetTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskRequestMultiError) AllErrors() []error { return m }

// GetTaskRequestValidationError is the validation error returned by
// GetTaskRequest.Validate if the designated constraints aren't met.
type GetTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskRequestValidationError) ErrorName() string { return "GetTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskRequestValidationError{}

// Validate checks the field values on GetTaskResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskResponseMultiError, or nil if none found.
func (m *GetTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTaskResponseMultiError(errors)
	}

	return nil
}

// GetTaskResponseMultiError is an error wrapping multiple validation errors
// returned by GetTaskResponse.ValidateAll() if the designated constraints
// aren't met.
type GetTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskResponseMultiError) AllErrors() []error { return m }

// GetTaskResponseValidationError is the validation error returned by
// GetTaskResponse.Validate if the designated constraints aren't met.
type GetTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskResponseValidationError) ErrorName() string { return "GetTaskResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskResponseValidationError{}
