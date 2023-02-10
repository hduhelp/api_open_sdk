package transfer

import (
	"fmt"
)

type ResponseError struct {
	Errors   []error // 请求错误
	HttpCode int     // http错误
	Code     int     // 业务错误码
	Msg      string  // 业务错误信息
	Remark   string  // 备注
}

func (r ResponseError) Error() string {
	var errStrings []string
	if r.HttpCode != 200 {
		errStrings = append(errStrings, fmt.Sprintf("http code: %d", r.HttpCode))
	}
	if len(r.Errors) != 0 {
		for _, err := range r.Errors {
			errStrings = append(errStrings, err.Error())
		}
	}
	if r.Code != 0 {
		errStrings = append(errStrings, fmt.Sprintf("code: %d, msg: %s", r.Code, r.Msg))
	}
	if r.Remark != "" {
		errStrings = append(errStrings, fmt.Sprintf("remark: %s", r.Remark))
	}
	return fmt.Sprintf("transfer error: %s", errStrings)
}
