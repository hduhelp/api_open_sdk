package staff

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type ErrInvalidStaffID struct {
	staffID string
}

func (e ErrInvalidStaffID) Error() string {
	return fmt.Sprintf("invalid staff_id: %s", e.staffID)
}

type StaffID struct {
	Raw    string
	Number string
	Unit   string
	Grade  string
	Type   Type
}

func (i StaffID) Valid() error {
	if _, err := ParseStaffID(i.Raw); err != nil {
		return ErrInvalidStaffID{i.Raw}
	}
	return nil
}

func ParseStaffID(str string) (StaffID, error) {
	var (
		unit  string
		grade string
		_type Type
	)
	pattern := regexp.MustCompile(`(\d+)`)
	matches := pattern.FindStringSubmatch(str)
	if len(matches) != 2 {
		return StaffID{Raw: str}, ErrInvalidStaffID{str}
	}
	switch len(matches[1]) {
	case 8:
		_type = Type_UNDERGRADUATE
		unit = matches[1][2:4]
		grade = "20" + matches[1][:2]
	case 9:
		_type = Type_POSTGRADUATE
		unit = matches[1][3:5]
		grade = "20" + matches[1][:2]
	case 5:
		_type = Type_TEACHER
	default:
		_type = Type_UNKNOWN
	}

	return StaffID{
		Raw:    str,
		Number: matches[1],
		Unit:   unit,
		Grade:  grade,
		Type:   _type,
	}, nil
}

func (i *StaffID) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("staffID cannot paste in string")
	}
	staffID, err := ParseStaffID(str)
	*i = staffID
	return err
}

func (i *StaffID) Value() (driver.Value, error) {
	return i.Raw, nil
}

func (i *StaffID) UnmarshalJSON(bytes []byte) error {
	var str string
	err := json.Unmarshal(bytes, &str)
	if err != nil {
		return err
	}
	staffID, err := ParseStaffID(str)
	if err != nil {
		return err
	}
	*i = staffID
	return nil
}

func (i StaffID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Raw)
}
