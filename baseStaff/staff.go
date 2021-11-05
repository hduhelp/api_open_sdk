package baseStaff

import (
	"github.com/gin-gonic/gin"
)

func Undergraduate(staffID string) *Staff {
	return &Staff{
		ID:   staffID,
		Type: Type_Undergraduate,
	}
}

func Teacher(staffID string) *Staff {
	return &Staff{
		ID:   staffID,
		Type: Type_Teacher,
	}
}

func Postgraduate(staffID string) *Staff {
	return &Staff{
		ID:   staffID,
		Type: Type_Postgraduate,
	}
}

func GetStaffFromGatewayGinContext(c *gin.Context) *Staff {
	return &Staff{
		ID:   GetStaffID(c),
		Type: GetStaffType(c),
	}
}

func GetStaffID(c *gin.Context) string {
	return c.GetString("staffId")
}

func GetStaffType(c *gin.Context) Type {
	switch c.GetString("staffType") {
	case "1":
		return Type_Undergraduate
	case "2":
		return Type_Teacher
	case "3":
		return Type_Postgraduate
	default:
		return Type_Unknown
	}
}

func GetSchool(c *gin.Context) string {
	return c.GetString("staffSchool")
}
