package person

import "github.com/hduhelp/api_open_sdk/baseStaff"

type InfoInterface interface {
	GetStaffID() string
	GetStaffName() string
	GetStaffType() baseStaff.Type
	GetStaffState() string
	GetUnitCode() string
	GetUnitName() string
}
