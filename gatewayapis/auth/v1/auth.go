package authv1

import "github.com/hduhelp/api_open_sdk/loh"

func (x *GetBindListByUserIdListResponse) UserIdList() []string {
	return loh.Array(x.Data, func(b *Bind) string {
		return b.UserId
	})
}

func (x *GetBindListByUserIdListResponse) UIDList() []string {
	return loh.Array(x.Data, func(b *Bind) string {
		return b.Uid
	})
}
