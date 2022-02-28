package test

import (
	"fmt"
	"github.com/hduhelp/api_open_sdk/baseStaff"
	"testing"
)

func TestNameStaffID(t *testing.T) {
	var staffID = "s19011141"
	data, err := baseStaff.ParseStaffID(staffID)
	fmt.Println(data, err)
}
