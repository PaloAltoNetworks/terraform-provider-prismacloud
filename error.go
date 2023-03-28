package prismacloud

import (
	"errors"
	"fmt"
	"strings"
)

var InvalidCredentialsError = errors.New("invalid credentials")
var ObjectNotFoundError = errors.New("object not found")
var AlreadyExistsError = errors.New("object already exists")
var InvalidPermissionGroupIdError = errors.New("invalid_permission_group_id") //permission group

type PrismaCloudErrorList struct {
	Errors     []PrismaCloudError
	Method     string
	StatusCode int
	Path       string
}

func (e PrismaCloudErrorList) Error() string {
	var buf strings.Builder
	buf.Grow(100)

	fmt.Fprintf(&buf, "%d/%s ", e.StatusCode, e.Path)
	for i := range e.Errors {
		if i != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(e.Errors[i].Error())
	}

	return buf.String()
}

func (e PrismaCloudErrorList) GenericError() error {
	for i := range e.Errors {
		if e.Errors[i].InvalidPermissionGroupIdError() {
			return InvalidPermissionGroupIdError
		} else if e.Errors[i].ObjectNotFound() {
			return ObjectNotFoundError
		} else if e.Errors[i].AlreadyExists() {
			return AlreadyExistsError
		}
	}

	return nil
}

type PrismaCloudError struct {
	Message  string `json:"i18nKey"`
	Severity string `json:"severity"`
	Subject  string `json:"subject"`
}

func (e PrismaCloudError) ObjectNotFound() bool {
	switch e.Message {
	case "invalid_id", "not_found":
		return true
	}

	return false
}
func (e PrismaCloudError) InvalidPermissionGroupIdError() bool {
	return strings.HasSuffix(e.Message, "invalid_permission_group_id")
}
func (e PrismaCloudError) AlreadyExists() bool {
	return strings.HasSuffix(e.Message, "_already_exists")
}

func (e PrismaCloudError) Error() string {
	return fmt.Sprintf("Error(msg:%s severity:%s subject:%v)", e.Message, e.Severity, e.Subject)
}

func (e PrismaCloudError) InternalError() bool {
	return strings.HasSuffix(e.Message, "internal_error")
}
func (e PrismaCloudError) AccountGroupNotFoundError() bool {
	return strings.HasSuffix(e.Message, "account_group_not_found")
}
