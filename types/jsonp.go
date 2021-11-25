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

	if err := json.Unmarshal(jsonpBytes, &j.Data); err != nil {
		return err
	}

	*j = Jsonp{
		JsonpJSON: render.JsonpJSON{
			Callback: callback,
			Data:     j.Data,
		},
	}
	return nil
}
