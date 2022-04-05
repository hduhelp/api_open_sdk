package test

import (
	"fmt"
	"testing"

	"github.com/hduhelp/api_open_sdk/campusapis/staff"
)

func TestNameStaffID(t *testing.T) {
	var staffID = "s19011141"
	data, err := staff.ParseStaffID(staffID)
	fmt.Println(data, err)
}
