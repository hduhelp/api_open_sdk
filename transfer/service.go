package transfer

var instance *Server

type Server struct {
	appKey   string
	appID    string
	endpoint string

	debug bool
}

func Init(appID, appKey string, options ...interface {
	apply()
}) {
	instance = defaultServer()
	instance.appID = appID
	instance.appKey = appKey
	for _, opt := range options {
		opt.apply()
	}
}

func (s *Server) Debug() *Server {
	s.debug = true
	return s
}

func defaultServer() *Server {
	return &Server{
		endpoint: "https://api.hduhelp.com/transfer",
	}
}
