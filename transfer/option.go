package transfer

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"net/url"
)

type option interface {
	apply()
}

func Endpoint(e string) OptionFunc {
	return func() {
		uri, err := url.ParseRequestURI(e)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "transfer: invalid endpoint url"))
		}
		conn, err := net.Dial("tcp", uri.Host)
		if err != nil {
			log.Fatalln("transfer: could not connect to server: ", err)
		}
		defer conn.Close()
		instance.endpoint = e
	}
}

type OptionFunc func()

func (o OptionFunc) apply() {
	o()
}
