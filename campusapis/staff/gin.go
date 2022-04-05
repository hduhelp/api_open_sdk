package staff

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
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

func first(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	return ss[0]
}

func FromMetadata(md metadata.MD) *Staff {
	staffID := first(md.Get("staffId"))
	if staff, ok := map[string]*Staff{
		"1": Undergraduate(staffID),
		"2": Postgraduate(staffID),
		"3": Teacher(staffID),
	}[first(md.Get("stafftype"))]; ok {
		return staff
	}
	return &Staff{ID: staffID}
}
