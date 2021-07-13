// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tracking/services/product.proto

package services

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ProductServiceGetSingleRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductServiceGetSingleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return ProductServiceGetSingleRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// ProductServiceGetSingleRequestValidationError is the validation error
// returned by ProductServiceGetSingleRequest.Validate if the designated
// constraints aren't met.
type ProductServiceGetSingleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductServiceGetSingleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductServiceGetSingleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductServiceGetSingleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductServiceGetSingleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductServiceGetSingleRequestValidationError) ErrorName() string {
	return "ProductServiceGetSingleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProductServiceGetSingleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductServiceGetSingleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductServiceGetSingleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductServiceGetSingleRequestValidationError{}

// Validate checks the field values on ProductServiceGetSingleResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductServiceGetSingleResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductServiceGetSingleResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProductServiceGetSingleResponseValidationError is the validation error
// returned by ProductServiceGetSingleResponse.Validate if the designated
// constraints aren't met.
type ProductServiceGetSingleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductServiceGetSingleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductServiceGetSingleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductServiceGetSingleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductServiceGetSingleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductServiceGetSingleResponseValidationError) ErrorName() string {
	return "ProductServiceGetSingleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductServiceGetSingleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductServiceGetSingleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductServiceGetSingleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductServiceGetSingleResponseValidationError{}

// Validate checks the field values on ProductServiceGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductServiceGetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetIds()) < 1 {
		return ProductServiceGetRequestValidationError{
			field:  "Ids",
			reason: "value must contain at least 1 item(s)",
		}
	}

	if m.GetPage() <= 0 {
		return ProductServiceGetRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
	}

	if m.GetLimit() <= 0 {
		return ProductServiceGetRequestValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// ProductServiceGetRequestValidationError is the validation error returned by
// ProductServiceGetRequest.Validate if the designated constraints aren't met.
type ProductServiceGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductServiceGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductServiceGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductServiceGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductServiceGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductServiceGetRequestValidationError) ErrorName() string {
	return "ProductServiceGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProductServiceGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductServiceGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductServiceGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductServiceGetRequestValidationError{}

// Validate checks the field values on ProductServiceGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductServiceGetResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProductServiceGetResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProductServiceGetResponseValidationError is the validation error returned by
// ProductServiceGetResponse.Validate if the designated constraints aren't met.
type ProductServiceGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductServiceGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductServiceGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductServiceGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductServiceGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductServiceGetResponseValidationError) ErrorName() string {
	return "ProductServiceGetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductServiceGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductServiceGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductServiceGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductServiceGetResponseValidationError{}

// Validate checks the field values on ProductServiceCreateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductServiceCreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductServiceCreateRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProductServiceCreateRequestValidationError is the validation error returned
// by ProductServiceCreateRequest.Validate if the designated constraints
// aren't met.
type ProductServiceCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductServiceCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductServiceCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductServiceCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductServiceCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductServiceCreateRequestValidationError) ErrorName() string {
	return "ProductServiceCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProductServiceCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductServiceCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductServiceCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductServiceCreateRequestValidationError{}

// Validate checks the field values on ProductServiceCreateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductServiceCreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductServiceCreateResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProductServiceCreateResponseValidationError is the validation error returned
// by ProductServiceCreateResponse.Validate if the designated constraints
// aren't met.
type ProductServiceCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductServiceCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductServiceCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductServiceCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductServiceCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductServiceCreateResponseValidationError) ErrorName() string {
	return "ProductServiceCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductServiceCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductServiceCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductServiceCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductServiceCreateResponseValidationError{}