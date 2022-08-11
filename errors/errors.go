// Package errors
// @Description: 异常处理包
package errors

import "bytes"

type BizError interface {
	error
	GetCode() string
	GetMsg() string
	GetErrors() *[]BizError
}

type SimpleBizError struct {
	Code   string      `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Errors *[]BizError `json:"errors,omitempty"`
}

type FieldError struct {
	*SimpleBizError
	Name string `json:"name,omitempty"`
}

func (err *SimpleBizError) Error() string {
	if err == nil {
		return ""
	}

	if err.Errors != nil {
		sb := bytes.Buffer{}
		sb.WriteString(err.Msg)

		for _, fe := range *err.Errors {
			sb.WriteByte('\n')
			sb.WriteString(fe.Error())
		}

		return sb.String()
	} else {
		return err.Msg
	}
}

func (err *SimpleBizError) GetCode() string {
	if err == nil {
		return ""
	}

	return err.Code
}

func (err *SimpleBizError) GetMsg() string {
	if err == nil {
		return ""
	}
	return err.Msg
}

func (err *SimpleBizError) AddError(cerr BizError) *SimpleBizError {
	if cerr == nil {
		return err
	}

	if err.Errors == nil {
		err.Errors = new([]BizError)
	}

	*err.Errors = append(*err.Errors, cerr)
	return err
}

func (err *SimpleBizError) GetErrors() *[]BizError {
	return err.Errors
}

func (err *SimpleBizError) HasError() bool {
	return err.Errors != nil && len(*err.Errors) > 0
}

const (
	CommonInvalidParams = "c.INVALID_PARAMS"
)

func InvalidParams() *SimpleBizError {
	return &SimpleBizError{Code: CommonInvalidParams}
}
