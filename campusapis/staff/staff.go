package staff

import (
	"github.com/gin-gonic/gin"
)

func Undergraduate(staffID string) *Staff {
	return &Staff{
		ID:   staffID,
		Type: Type_UNDERGRADUATE,
	}
}

func Teacher(staffID string) *Staff {
	return &Staff{
		ID:   staffID,
		Type: Type_TEACHER,
	}
}

func Postgraduate(staffID string) *Staff {
	return &Staff{
		ID:   staffID,
		Type: Type_POSTGRADUATE,
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
		return Type_UNDERGRADUATE
	case "2":
		return Type_TEACHER
	case "3":
		return Type_POSTGRADUATE
	default:
		return Type_UNKNOWN
	}
}

func GetSchool(c *gin.Context) string {
	return c.GetString("staffSchool")
}
