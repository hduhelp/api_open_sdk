package baseStaff

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FromContext(c *gin.Context) *Staff {
	staffID := c.GetHeader("staffId")
	if staff, ok := map[string]*Staff{
		"1": Undergraduate(staffID),
		"2": Postgraduate(staffID),
		"3": Teacher(staffID),
	}[c.GetHeader("staffType")]; ok {
		return staff
	}

	return c.MustGet("staff").(*Staff)
}

func FromHeader(header http.Header) *Staff {
	staffID := header.Get("staffId")
	if staff, ok := map[string]*Staff{
		"1": Undergraduate(staffID),
		"2": Postgraduate(staffID),
		"3": Teacher(staffID),
	}[header.Get("staffType")]; ok {
		return staff
	}

	return &Staff{ID: staffID}
}
