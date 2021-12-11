package transfer

import "fmt"

type ResponseError struct {
	string
	int
}

func (r ResponseError) Error() string {
	return fmt.Sprintf("transfer: code(%d) %s", r.int, r.string)
}
