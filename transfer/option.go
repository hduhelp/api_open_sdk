package transfer

import (
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"net"
	"net/url"
)

func Endpoint(e string) ServiceOptionFunc {
	return func(s *Server) {
		uri, err := url.ParseRequestURI(e)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "transfer: invalid endpoint url"))
		}
		conn, err := net.Dial("tcp", uri.Host)
		if err != nil {
			log.Fatalln("transfer: could not connect to server: ", err)
		}
		defer conn.Close()
		s.endpoint = e
	}
}

func Debug() ServiceOptionFunc {
	return func(s *Server) {
		s.debug = true
	}
}

type ServiceOptionFunc func(*Server)

func (o ServiceOptionFunc) apply() {
	o(instance)
}

type School struct {
	string
}

func (s School) apply(request *Request) {
	request.SuperAgent.Set("x-hduhelp-school", s.string)
}

type Cache struct {
	bool
}

func (s Cache) apply(request *Request) {
	if !s.bool {
		request.SuperAgent.Header.Del("x-hduhelp-cache")
	}
}

type PropagatorManager struct {
	propagator propagation.TextMapPropagator
}

func WithPropagator(propagator propagation.TextMapPropagator) interface{ apply() } {
	return PropagatorManager{propagator: propagator}
}

func (p PropagatorManager) apply() {
	instance.propagator = p.propagator
}
