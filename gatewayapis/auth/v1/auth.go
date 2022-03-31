package authv1

import "github.com/hduhelp/api_open_sdk/loh"

func (x *BindListResponse) UserIdList() []string {
	return loh.Array(x.Items, func(b *Bind) string {
		return b.UserId
	})
}

func (x *BindListResponse) UIDList() []string {
	return loh.Array(x.Items, func(b *Bind) string {
		return b.Uid
	})
}
