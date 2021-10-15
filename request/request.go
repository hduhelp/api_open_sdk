package request

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
)

type Request struct {
	token string
	*ResponseData
	*http.Response
	Errors []error
	*gorequest.SuperAgent
}

func New() *Request {
	return &Request{
		SuperAgent: gorequest.New(),
	}
}

func (r *Request) SetToken(token string) *Request {
	r.token = token
	r.SuperAgent.Set("Authorization", "token "+token)
	return r
}

func (r *Request) EndData(u interface{}) {
	r.ResponseData.Data = u
	r.Response, _, r.Errors = r.SuperAgent.EndStruct(r.ResponseData)

}
