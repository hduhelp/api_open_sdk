package transfer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hduhelp/api_open_sdk/utils"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Response struct {
	Cache bool        `json:"cache"`
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
}

type Request struct {
	*gorequest.SuperAgent
	ResponseData *Response
	*http.Response

	*body
	url.Values

	done bool
	error
	code int
}

func Get(appID string, path string, param map[string]string, staffID string) (resp *Request) {
	return NewRequest(instance.appID, appID, path, param, staffID, "GET", nil)
}

func Post(appID string, path string, param map[string]string, staffID string, postForm interface{}) (resp *Request) {
	return NewRequest(instance.appID, appID, path, param, staffID, "POST", postForm)
}

func Put(appID string, path string, param map[string]string, staffID string, postForm interface{}) (resp *Request) {
	return NewRequest(instance.appID, appID, path, param, staffID, "PUT", postForm)
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

func NewRequest(from string, to string, path string, param map[string]string, staffID string, method string, postBody interface{}) *Request {
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
	return request
}

func (r *Request) make() {
	dataBodyB, err := json.Marshal(r.body)
	if err != nil {
		r.doneWithError(40301, err)
		return
	}

	queries := map[string]string{
		"from":      r.From,
		"to":        r.To,
		"path":      r.Path,
		"timestamp": r.timestampStr(),
		"body":      base64.StdEncoding.EncodeToString(dataBodyB),
	}

	r.setToken()
	r.SuperAgent.
		Post(instance.endpoint).
		Set("sign", "sign "+r.sign()).
		Set("x-hduhelp-cache", "no-cache").
		Query(queries).
		Send(r.body)
	return
}

func (r *Request) doneWithError(code int, err error) {
	r.error = err
	r.code = code
	r.done = true
}

func (r *Request) EndStruct(data interface{}) (int, int, error) {
	r.make()
	r.ResponseData = new(Response)
	r.ResponseData.Data = data
	if !r.done {
		var errs []error
		r.Response, _, errs = r.SuperAgent.EndStruct(&r.ResponseData)
		if len(errs) != 0 {
			r.doneWithError(50000, errs[0])
			if r.Response == nil {
				return 50000, 500, r.error
			}
			return 50000, r.Response.StatusCode, errs[0]
		}
	}
	return r.ResponseData.Error, r.Response.StatusCode, nil
}

func (r *Request) End() (*Response, error) {
	var in interface{}
	_, _, err := r.EndStruct(&in)
	return r.ResponseData, err
}

func (r Request) timestampStr() string {
	return strconv.FormatInt(r.Timestamp, 10)
}

func (r *Request) sign() string {
	return utils.Sha1Encode(fmt.Sprintf("%s%s%s%s", r.From, r.To, instance.appKey, r.timestampStr()))
}
