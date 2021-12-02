package transfer

import "os"

func (r *Request) setToken() {
	token := os.Getenv("HDUHELP_TOKEN")
	if token != "" {
		r.SuperAgent.Set("Authorization", "token "+token)
	}
}
