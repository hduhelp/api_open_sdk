package baseStaff

type Type uint

const (
	TypeStudent Type = iota
	TypeTeacher
)

type Staff struct {
	ID   string
	Type Type
}

func Teacher(staffID string) Staff {
	return Staff{
		ID:   staffID,
		Type: TypeTeacher,
	}
}

func Student(staffID string) Staff {
	return Staff{
		ID:   staffID,
		Type: TypeStudent,
	}
}
