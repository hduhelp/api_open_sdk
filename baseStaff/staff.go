package baseStaff

func UndergraduateStudent(staffID string) Staff {
	return Staff{
		ID:   staffID,
		Type: Type_Undergraduate,
	}
}

func Teacher(staffID string) Staff {
	return Staff{
		ID:   staffID,
		Type: Type_Teacher,
	}
}

func PostgraduateStudent(staffID string) Staff {
	return Staff{
		ID:   staffID,
		Type: Type_Postgraduate,
	}
}
