package baseStaff

import "github.com/gin-gonic/gin"

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
