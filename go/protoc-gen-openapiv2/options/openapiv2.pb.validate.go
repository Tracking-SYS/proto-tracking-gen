// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protoc-gen-openapiv2/options/openapiv2.proto

package options

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

// Validate checks the field values on Swagger with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Swagger) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Swagger

	if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SwaggerValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Host

	// no validation rules for BasePath

	for key, val := range m.GetResponses() {
		_ = val

		// no validation rules for Responses[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SwaggerValidationError{
					field:  fmt.Sprintf("Responses[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetSecurityDefinitions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SwaggerValidationError{
				field:  "SecurityDefinitions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSecurity() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SwaggerValidationError{
					field:  fmt.Sprintf("Security[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetExternalDocs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SwaggerValidationError{
				field:  "ExternalDocs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetExtensions() {
		_ = val

		// no validation rules for Extensions[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SwaggerValidationError{
					field:  fmt.Sprintf("Extensions[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SwaggerValidationError is the validation error returned by Swagger.Validate
// if the designated constraints aren't met.
type SwaggerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SwaggerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SwaggerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SwaggerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SwaggerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SwaggerValidationError) ErrorName() string { return "SwaggerValidationError" }

// Error satisfies the builtin error interface
func (e SwaggerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSwagger.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SwaggerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SwaggerValidationError{}

// Validate checks the field values on Operation with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Operation) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Summary

	// no validation rules for Description

	if v, ok := interface{}(m.GetExternalDocs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OperationValidationError{
				field:  "ExternalDocs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperationId

	for key, val := range m.GetResponses() {
		_ = val

		// no validation rules for Responses[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OperationValidationError{
					field:  fmt.Sprintf("Responses[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Deprecated

	for idx, item := range m.GetSecurity() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OperationValidationError{
					field:  fmt.Sprintf("Security[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for key, val := range m.GetExtensions() {
		_ = val

		// no validation rules for Extensions[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OperationValidationError{
					field:  fmt.Sprintf("Extensions[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OperationValidationError is the validation error returned by
// Operation.Validate if the designated constraints aren't met.
type OperationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OperationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OperationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OperationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OperationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OperationValidationError) ErrorName() string { return "OperationValidationError" }

// Error satisfies the builtin error interface
func (e OperationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOperation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OperationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OperationValidationError{}

// Validate checks the field values on Header with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Header) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	// no validation rules for Type

	// no validation rules for Format

	// no validation rules for Default

	// no validation rules for Pattern

	return nil
}

// HeaderValidationError is the validation error returned by Header.Validate if
// the designated constraints aren't met.
type HeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValidationError) ErrorName() string { return "HeaderValidationError" }

// Error satisfies the builtin error interface
func (e HeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValidationError{}

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	if v, ok := interface{}(m.GetSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResponseValidationError{
				field:  "Schema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetHeaders() {
		_ = val

		// no validation rules for Headers[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  fmt.Sprintf("Headers[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Examples

	for key, val := range m.GetExtensions() {
		_ = val

		// no validation rules for Extensions[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  fmt.Sprintf("Extensions[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on Info with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Info) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for TermsOfService

	if v, ok := interface{}(m.GetContact()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoValidationError{
				field:  "Contact",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLicense()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoValidationError{
				field:  "License",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Version

	for key, val := range m.GetExtensions() {
		_ = val

		// no validation rules for Extensions[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InfoValidationError{
					field:  fmt.Sprintf("Extensions[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// InfoValidationError is the validation error returned by Info.Validate if the
// designated constraints aren't met.
type InfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoValidationError) ErrorName() string { return "InfoValidationError" }

// Error satisfies the builtin error interface
func (e InfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoValidationError{}

// Validate checks the field values on Contact with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Contact) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Url

	// no validation rules for Email

	return nil
}

// ContactValidationError is the validation error returned by Contact.Validate
// if the designated constraints aren't met.
type ContactValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContactValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContactValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContactValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContactValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContactValidationError) ErrorName() string { return "ContactValidationError" }

// Error satisfies the builtin error interface
func (e ContactValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContact.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContactValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContactValidationError{}

// Validate checks the field values on License with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *License) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Url

	return nil
}

// LicenseValidationError is the validation error returned by License.Validate
// if the designated constraints aren't met.
type LicenseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LicenseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LicenseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LicenseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LicenseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LicenseValidationError) ErrorName() string { return "LicenseValidationError" }

// Error satisfies the builtin error interface
func (e LicenseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLicense.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LicenseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LicenseValidationError{}

// Validate checks the field values on ExternalDocumentation with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExternalDocumentation) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	// no validation rules for Url

	return nil
}

// ExternalDocumentationValidationError is the validation error returned by
// ExternalDocumentation.Validate if the designated constraints aren't met.
type ExternalDocumentationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExternalDocumentationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExternalDocumentationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExternalDocumentationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExternalDocumentationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExternalDocumentationValidationError) ErrorName() string {
	return "ExternalDocumentationValidationError"
}

// Error satisfies the builtin error interface
func (e ExternalDocumentationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExternalDocumentation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExternalDocumentationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExternalDocumentationValidationError{}

// Validate checks the field values on Schema with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Schema) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetJsonSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaValidationError{
				field:  "JsonSchema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Discriminator

	// no validation rules for ReadOnly

	if v, ok := interface{}(m.GetExternalDocs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaValidationError{
				field:  "ExternalDocs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Example

	return nil
}

// SchemaValidationError is the validation error returned by Schema.Validate if
// the designated constraints aren't met.
type SchemaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaValidationError) ErrorName() string { return "SchemaValidationError" }

// Error satisfies the builtin error interface
func (e SchemaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchema.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaValidationError{}

// Validate checks the field values on JSONSchema with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JSONSchema) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Ref

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for Default

	// no validation rules for ReadOnly

	// no validation rules for Example

	// no validation rules for MultipleOf

	// no validation rules for Maximum

	// no validation rules for ExclusiveMaximum

	// no validation rules for Minimum

	// no validation rules for ExclusiveMinimum

	// no validation rules for MaxLength

	// no validation rules for MinLength

	// no validation rules for Pattern

	// no validation rules for MaxItems

	// no validation rules for MinItems

	// no validation rules for UniqueItems

	// no validation rules for MaxProperties

	// no validation rules for MinProperties

	// no validation rules for Format

	return nil
}

// JSONSchemaValidationError is the validation error returned by
// JSONSchema.Validate if the designated constraints aren't met.
type JSONSchemaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JSONSchemaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JSONSchemaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JSONSchemaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JSONSchemaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JSONSchemaValidationError) ErrorName() string { return "JSONSchemaValidationError" }

// Error satisfies the builtin error interface
func (e JSONSchemaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJSONSchema.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JSONSchemaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JSONSchemaValidationError{}

// Validate checks the field values on Tag with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Tag) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	if v, ok := interface{}(m.GetExternalDocs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TagValidationError{
				field:  "ExternalDocs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TagValidationError is the validation error returned by Tag.Validate if the
// designated constraints aren't met.
type TagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagValidationError) ErrorName() string { return "TagValidationError" }

// Error satisfies the builtin error interface
func (e TagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagValidationError{}

// Validate checks the field values on SecurityDefinitions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SecurityDefinitions) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetSecurity() {
		_ = val

		// no validation rules for Security[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecurityDefinitionsValidationError{
					field:  fmt.Sprintf("Security[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SecurityDefinitionsValidationError is the validation error returned by
// SecurityDefinitions.Validate if the designated constraints aren't met.
type SecurityDefinitionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecurityDefinitionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecurityDefinitionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecurityDefinitionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecurityDefinitionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecurityDefinitionsValidationError) ErrorName() string {
	return "SecurityDefinitionsValidationError"
}

// Error satisfies the builtin error interface
func (e SecurityDefinitionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecurityDefinitions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecurityDefinitionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecurityDefinitionsValidationError{}

// Validate checks the field values on SecurityScheme with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SecurityScheme) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	// no validation rules for Description

	// no validation rules for Name

	// no validation rules for In

	// no validation rules for Flow

	// no validation rules for AuthorizationUrl

	// no validation rules for TokenUrl

	if v, ok := interface{}(m.GetScopes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SecuritySchemeValidationError{
				field:  "Scopes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetExtensions() {
		_ = val

		// no validation rules for Extensions[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecuritySchemeValidationError{
					field:  fmt.Sprintf("Extensions[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SecuritySchemeValidationError is the validation error returned by
// SecurityScheme.Validate if the designated constraints aren't met.
type SecuritySchemeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecuritySchemeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecuritySchemeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecuritySchemeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecuritySchemeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecuritySchemeValidationError) ErrorName() string { return "SecuritySchemeValidationError" }

// Error satisfies the builtin error interface
func (e SecuritySchemeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecurityScheme.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecuritySchemeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecuritySchemeValidationError{}

// Validate checks the field values on SecurityRequirement with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SecurityRequirement) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetSecurityRequirement() {
		_ = val

		// no validation rules for SecurityRequirement[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecurityRequirementValidationError{
					field:  fmt.Sprintf("SecurityRequirement[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SecurityRequirementValidationError is the validation error returned by
// SecurityRequirement.Validate if the designated constraints aren't met.
type SecurityRequirementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecurityRequirementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecurityRequirementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecurityRequirementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecurityRequirementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecurityRequirementValidationError) ErrorName() string {
	return "SecurityRequirementValidationError"
}

// Error satisfies the builtin error interface
func (e SecurityRequirementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecurityRequirement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecurityRequirementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecurityRequirementValidationError{}

// Validate checks the field values on Scopes with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Scopes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Scope

	return nil
}

// ScopesValidationError is the validation error returned by Scopes.Validate if
// the designated constraints aren't met.
type ScopesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScopesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScopesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScopesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScopesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScopesValidationError) ErrorName() string { return "ScopesValidationError" }

// Error satisfies the builtin error interface
func (e ScopesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScopes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScopesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScopesValidationError{}

// Validate checks the field values on
// SecurityRequirement_SecurityRequirementValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SecurityRequirement_SecurityRequirementValue) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SecurityRequirement_SecurityRequirementValueValidationError is the
// validation error returned by
// SecurityRequirement_SecurityRequirementValue.Validate if the designated
// constraints aren't met.
type SecurityRequirement_SecurityRequirementValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecurityRequirement_SecurityRequirementValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecurityRequirement_SecurityRequirementValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecurityRequirement_SecurityRequirementValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecurityRequirement_SecurityRequirementValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecurityRequirement_SecurityRequirementValueValidationError) ErrorName() string {
	return "SecurityRequirement_SecurityRequirementValueValidationError"
}

// Error satisfies the builtin error interface
func (e SecurityRequirement_SecurityRequirementValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecurityRequirement_SecurityRequirementValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecurityRequirement_SecurityRequirementValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecurityRequirement_SecurityRequirementValueValidationError{}
