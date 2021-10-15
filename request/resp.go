package request

type ResponseData struct {
	Cache bool        `json:"cache"`
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
}
