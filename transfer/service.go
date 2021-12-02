package transfer

var instance *server

type server struct {
	appKey   string
	appID    string
	endpoint string
}

func Init(appID, appKey string, options ...option) {
	instance = defaultServer()
	instance.appID = appID
	instance.appKey = appKey
	for _, opt := range options {
		opt.apply()
	}
}

func defaultServer() *server {
	return &server{
		endpoint: "https://api.hduhelp.com/transfer",
	}
}
