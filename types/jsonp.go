package types

import (
	"encoding/json"
	"github.com/gin-gonic/gin/render"
)

type Jsonp struct {
	render.JsonpJSON
}

func (j *Jsonp) UnmarshalJSON(bytes []byte) error {
	var callback string
	var callbackIndex int
	var jsonpBytes []byte
	var data interface{}

	for i, v := range bytes {
		if v == '(' {
			callback = string(bytes[:i])
			callbackIndex = i
		}
		if v == ')' {
			jsonpBytes = bytes[callbackIndex+1 : i]
			break
		}
	}

	if err := json.Unmarshal(jsonpBytes, &data); err != nil {
		return err
	}

	*j = Jsonp{
		JsonpJSON: render.JsonpJSON{
			Callback: callback,
			Data:     data,
		},
	}
	return nil
}
