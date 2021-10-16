package baseStaff

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
