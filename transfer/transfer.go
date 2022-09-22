package transfer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hduhelp/api_open_sdk/utils"
	"github.com/parnurzeal/gorequest"
)

type Response struct {
	Cache bool        `json:"cache"`
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
}

func (r *Response) UnmarshalJSON(bytes []byte) error {
	d := struct {
		Cache bool        `json:"cache"`
		Data  interface{} `json:"data"`
		Error int         `json:"error"`
		Msg   string      `json:"msg"`
	}{
		Data: r.Data,
	}
	if err := json.Unmarshal(bytes, &d); err == nil {
		*r = d
		return err
	}
	d.Data = nil
	err := json.Unmarshal(bytes, &d)
	*r = d
	return err
}

func (r Response) error() error {
	if r.Error == 0 {
		return nil
	}
	return &ResponseError{
		string: r.Msg,
		int:    r.Error,
	}
}

type Request struct {
	*gorequest.SuperAgent
	ResponseData *Response
	RawData      string

	*http.Response

	*body
	url.Values

	done bool
	error
	code int

	options []interface {
		apply(*Request)
	}
}

func Get(appID string, path string, param map[string]string, staffID string, options ...interface {
	apply(*Request)
}) (resp *Request) {
	return NewRequest(instance.appID, appID, path, param, staffID, http.MethodGet, nil, options...)
}

func Post(appID string, path string, param map[string]string, staffID string, postForm interface{}, options ...interface {
	apply(*Request)
}) (resp *Request) {
	return NewRequest(instance.appID, appID, path, param, staffID, http.MethodPost, postForm, options...)
}

func Put(appID string, path string, param map[string]string, staffID string, postForm interface{}, options ...interface {
	apply(*Request)
}) (resp *Request) {
	return NewRequest(instance.appID, appID, path, param, staffID, http.MethodPut, postForm, options...)
}

type body struct {
	From      string            `json:"from"`
	To        string            `json:"to"`
	Timestamp int64             `json:"timestamp"`
	Path      string            `json:"path"`
	Query     map[string]string `json:"query"`
	StaffID   string            `json:"staffID"`
	Method    string            `json:"method"`
	PostForm  interface{}       `json:"postForm"`
}

func NewRequest(from string, to string, path string, param map[string]string,
	staffID string, method string, postBody interface{}, options ...interface {
		apply(*Request)
	}) *Request {
	request := &Request{
		SuperAgent: gorequest.New(),
	}

	timestamp := time.Now().Unix()
	request.body = &body{
		From:      from,
		To:        to,
		Timestamp: timestamp,
		Path:      path,
		Query:     param,
		StaffID:   staffID,
		Method:    method,
		PostForm:  postBody,
	}
	request.options = options
	return request
}

func (r *Request) make() {
	r.SuperAgent.
		Post(instance.endpoint).
		Set("sign", "sign "+r.sign()).
		Set("x-hduhelp-cache", "no-cache").
		Send(r.body)
	for _, option := range r.options {
		option.apply(r)
	}
	r.setToken()
}

func (r *Request) doneWithError(code int, err error) {
	r.error = err
	r.code = code
	r.done = true
}

func (r *Request) AddHeader(key, value string) *Request {
	r.SuperAgent.Set(key, value)
	return r
}

func (r *Request) EndStruct(data interface{}) (int, int, error) {
	var bytes []byte
	r.make()
	r.ResponseData = new(Response)
	r.ResponseData.Data = data
	if !r.done {
		var errs []error
		r.Response, bytes, errs = r.SuperAgent.EndStruct(&r.ResponseData)
		r.RawData = string(bytes)
		if len(errs) != 0 {
			r.doneWithError(50000, errs[0])
			if r.Response == nil {
				return 50000, 500, r.error
			}
			return 50000, r.Response.StatusCode, errs[0]
		}
	}
	return r.ResponseData.Error, r.Response.StatusCode, r.ResponseData.error()
}

func (r *Request) End() (*Response, error) {
	var in interface{}
	_, _, err := r.EndStruct(in)
	return r.ResponseData, err
}

func (r Request) timestampStr() string {
	return strconv.FormatInt(r.Timestamp, 10)
}

func (r *Request) sign() string {
	return utils.Sha1Encode(fmt.Sprintf("%s%s%s%s", r.From, r.To, instance.appKey, r.timestampStr()))
}
