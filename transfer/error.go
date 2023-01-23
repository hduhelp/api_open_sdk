package transfer

import "fmt"

type ResponseError struct {
	HttpCode  int     // http错误
	HttpError []error // http错误
	Code      int     // 业务错误
	Msg       string  // 业务错误
}

func (r ResponseError) Error() string {
	var errStrings []string
	if r.HttpCode != 200 {
		errStrings = append(errStrings, fmt.Sprintf("http code: %d", r.HttpCode))
	}
	if len(r.HttpError) != 0 {
		for _, err := range r.HttpError {
			errStrings = append(errStrings, err.Error())
		}
	}
	if r.Code != 0 {
		errStrings = append(errStrings, fmt.Sprintf("code: %d, msg: %s", r.Code, r.Msg))
	}
	return fmt.Sprintf("transfer error: %s", errStrings)
}
