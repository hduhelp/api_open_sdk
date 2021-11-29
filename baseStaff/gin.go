package baseStaff

import "github.com/gin-gonic/gin"

func FromContext(c *gin.Context) *Staff {
	staffID := c.GetHeader("staffId")
	if staff, ok := map[string]*Staff{
		"1": Undergraduate(staffID),
		"2": Teacher(staffID),
		"3": Postgraduate(staffID),
	}[c.GetHeader("staffType")]; ok {
		return staff
	}

	if staff, ok := map[string]*Staff{
		"10": Undergraduate(staffID),
		"11": Undergraduate(staffID),
		"30": Teacher(staffID),
		"31": Teacher(staffID),
		"20": Postgraduate(staffID),
		"21": Postgraduate(staffID),

		"":   {ID: staffID, Type: Type_Unknown},
		"90": {ID: staffID, Type: Type_Unknown},
	}[c.GetHeader("userType")]; ok {
		return staff
	}

	return c.MustGet("staff").(*Staff)
}
