// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authzed/api/v1/watch_service.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on WatchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WatchRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetOptionalObjectTypes() {
		_, _ = idx, item

		if len(item) > 128 {
			return WatchRequestValidationError{
				field:  fmt.Sprintf("OptionalObjectTypes[%v]", idx),
				reason: "value length must be at most 128 bytes",
			}
		}

		if !_WatchRequest_OptionalObjectTypes_Pattern.MatchString(item) {
			return WatchRequestValidationError{
				field:  fmt.Sprintf("OptionalObjectTypes[%v]", idx),
				reason: "value does not match regex pattern \"^([a-z][a-z0-9_]{1,62}[a-z0-9]/)?[a-z][a-z0-9_]{1,62}[a-z0-9]$\"",
			}
		}

	}

	if v, ok := interface{}(m.GetOptionalStartCursor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WatchRequestValidationError{
				field:  "OptionalStartCursor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WatchRequestValidationError is the validation error returned by
// WatchRequest.Validate if the designated constraints aren't met.
type WatchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchRequestValidationError) ErrorName() string { return "WatchRequestValidationError" }

// Error satisfies the builtin error interface
func (e WatchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchRequestValidationError{}

var _WatchRequest_OptionalObjectTypes_Pattern = regexp.MustCompile("^([a-z][a-z0-9_]{1,62}[a-z0-9]/)?[a-z][a-z0-9_]{1,62}[a-z0-9]$")

// Validate checks the field values on WatchResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WatchResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUpdates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WatchResponseValidationError{
					field:  fmt.Sprintf("Updates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetChangesThrough()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WatchResponseValidationError{
				field:  "ChangesThrough",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WatchResponseValidationError is the validation error returned by
// WatchResponse.Validate if the designated constraints aren't met.
type WatchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchResponseValidationError) ErrorName() string { return "WatchResponseValidationError" }

// Error satisfies the builtin error interface
func (e WatchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchResponseValidationError{}
