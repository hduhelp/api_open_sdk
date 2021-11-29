package person

import "github.com/hduhelp/api_open_sdk/baseStaff"

type InfoInterface interface {
	GetPersonInfo(info *baseStaff.Staff) (*Info, error)
}
