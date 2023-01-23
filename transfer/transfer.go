package transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hduhelp/api_open_sdk/utils"
	"github.com/parnurzeal/gorequest"
)

const tracerName = "github.com/hduhelp/api_open_sdk/transfer"

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

type Request struct {
	*gorequest.SuperAgent
	ResponseData *Response
	RawData      string

	*http.Response

	*body
	url.Values

	ctx context.Context

	options []interface {
		apply(*Request)
	}
}

func Get(ctx context.Context, appID string, path string, param map[string]string, staffID string, options ...interface {
	apply(*Request)
}) (resp *Request) {
	return NewRequest(ctx, instance.appID, appID, path, param, staffID, http.MethodGet, nil, options...)
}

func Post(ctx context.Context, appID string, path string, param map[string]string, staffID string, postForm interface{}, options ...interface {
	apply(*Request)
}) (resp *Request) {
	return NewRequest(ctx, instance.appID, appID, path, param, staffID, http.MethodPost, postForm, options...)
}

func Put(ctx context.Context, appID string, path string, param map[string]string, staffID string, postForm interface{}, options ...interface {
	apply(*Request)
}) (resp *Request) {
	return NewRequest(ctx, instance.appID, appID, path, param, staffID, http.MethodPut, postForm, options...)
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

func NewRequest(ctx context.Context, from string, to string, path string, param map[string]string,
	staffID string, method string, postBody interface{}, options ...interface {
		apply(*Request)
	}) *Request {
	request := &Request{
		SuperAgent: gorequest.New(),
		ctx:        ctx,
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
		Post(instance.endpoint). // Post() 会把 SuperAgent 清空
		Set("sign", "sign "+r.sign()).
		Set("x-hduhelp-cache", "no-cache").
		Send(r.body)
	for _, option := range r.options {
		option.apply(r)
	}
	r.setToken()
}

func (r *Request) AddHeader(key, value string) *Request {
	r.SuperAgent.Set(key, value)
	return r
}

func (r *Request) EndStruct(data interface{}) error {
	opts := []trace.SpanStartOption{
		trace.WithAttributes(r.AttributesFromRequestBody()...),
	}
	newCtx, span := otel.Tracer(tracerName).Start(r.ctx, fmt.Sprintf("%s-%s-%s", r.body.Method, r.body.To, r.body.Path), opts...)
	defer span.End()

	r.make() // 会把 SuperAgent 清空
	instance.getPropagator().Inject(newCtx, propagation.HeaderCarrier(r.SuperAgent.Header))

	var bytes []byte
	r.ResponseData = new(Response)
	r.ResponseData.Data = data
	var errs []error
	r.Response, bytes, errs = r.SuperAgent.EndStruct(&r.ResponseData)
	r.RawData = string(bytes)
	if len(errs) != 0 {
		return &ResponseError{
			HttpCode:  r.Response.StatusCode,
			HttpError: errs,
		}
	}
	if r.ResponseData.Error != 0 {
		return &ResponseError{
			HttpCode:  r.Response.StatusCode,
			HttpError: errs,
			Code:      r.ResponseData.Error,
			Msg:       r.ResponseData.Msg,
		}
	}
	return nil
}

func (r *Request) End() (*Response, error) {
	var in interface{}
	err := r.EndStruct(in)
	return r.ResponseData, err
}

func (r Request) timestampStr() string {
	return strconv.FormatInt(r.Timestamp, 10)
}

func (r *Request) sign() string {
	return utils.Sha1Encode(fmt.Sprintf("%s%s%s%s", r.From, r.To, instance.appKey, r.timestampStr()))
}

func (r *Request) AttributesFromRequestBody() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("from", r.body.From),
		attribute.String("to", r.body.To),
		attribute.Int64("timestamp", r.body.Timestamp),
		attribute.String("path", r.body.Path),
		attribute.String("staffID", r.body.StaffID),
		attribute.String("method", r.body.Method),
	}
}
