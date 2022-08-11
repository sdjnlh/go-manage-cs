// Package web
// @Description: 程序接口定义包，方法返回消息封装，接口注册
package web

import (
	"bytes"
	"encoding/json"
	"jinbao-cs/errors"
	"strconv"
)

type IResult interface {
	IsOk() bool
	Err() errors.BizError
	SetCode(code string)
	SetMsg(msg string)
	SetError(err errors.BizError)
	Set(key string, value interface{})
}

type Result struct {
	Ok    bool                   `json:"ok"`
	Code  string                 `json:"code"`
	Msg   string                 `json:"msg"`
	Error errors.BizError        `json:"err,omitempty"`
	Data  interface{}            `json:"data,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty"`
}

func (r *Result) IsOk() bool {
	return r.Ok
}

func (r *Result) Set(key string, value interface{}) {
	if r.Extra == nil {
		r.Extra = map[string]interface{}{}
	}
	r.Extra[key] = value
}

func (r *Result) SetCode(code string) {
	r.Code = code
}

func (r *Result) SetMsg(msg string) {
	r.Msg = msg
}

func (r *Result) Err() errors.BizError {
	return r.Error
}

func (r *Result) SetError(err errors.BizError) {
	r.Error = err
}

func (r *Result) Failure(errs ...errors.BizError) *Result {
	r.Ok = false
	if len(errs) > 0 {
		r.Error = errs[0]
	}
	return r
}

func (r *Result) FailureWithData(data interface{}, err errors.BizError) *Result {
	r.Ok = false
	r.Error = err
	r.Data = data

	return r
}

func (r *Result) Success(ds ...interface{}) *Result {
	r.Ok = true
	if len(ds) > 0 {
		r.Data = ds[0]
	}
	return r
}

func NewResult(data interface{}) *Result {
	return &Result{Error: &errors.SimpleBizError{}, Data: data}
}

type FilterResult struct {
	Result
}

func NewFilterResult(data interface{}) *FilterResult {
	return &FilterResult{
		Result: Result{Error: &errors.SimpleBizError{}, Data: data},
	}
}

type StringArray []string

func serializeStringArray(s []string, prefix string, suffix string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)

	for index, str := range s {
		if index > 0 {
			buffer.WriteByte(',')
		}
		buffer.WriteByte('\'')
		buffer.WriteString(str)
		buffer.WriteByte('\'')
	}
	buffer.WriteString(suffix)
	return buffer.Bytes()
}
func (s *StringArray) ToIn() string {
	return string(serializeStringArray(*s, "(", ")"))
}

type Int64Array []int64

func (s *Int64Array) Contains(val int64) bool {
	if s == nil {
		return false
	}

	for _, v := range *s {
		if v == val {
			return true
		}
	}

	return false
}

func serializeBigIntArray(s []int64, prefix string, suffix string) []byte {
	var buffer bytes.Buffer

	buffer.WriteString(prefix)

	for idx, val := range s {
		if idx > 0 {
			buffer.WriteByte(',')
		}
		buffer.WriteString(strconv.FormatInt(val, 10))
	}

	buffer.WriteString(suffix)

	return buffer.Bytes()
}

func (s *Int64Array) ToIn() string {
	return string(serializeBigIntArray(*s, "(", ")"))
}

type Context map[string]interface{}

func (ctx *Context) MustGet(key string) interface{} {
	v := (*ctx)[key]

	if v == nil {
		panic("key " + key + " not present in context")
	}
	return v
}

func (ctx *Context) Get(key string) interface{} {
	return (*ctx)[key]
}

func (ctx *Context) Set(key string, value interface{}) {
	(*ctx)[key] = value
}

type JsonMap map[string]interface{}

func (c *JsonMap) FromDB(bytes []byte) error {
	return json.Unmarshal(bytes, c)
}

func (c *JsonMap) ToDB() (bytes []byte, err error) {
	if c == nil {
		return []byte("{}"), nil
	}
	bytes, err = json.Marshal(c)
	return
}

func (c *JsonMap) String() string {
	if c == nil {
		return ""
	}
	bts, _ := json.Marshal(c)
	return string(bts)
}
