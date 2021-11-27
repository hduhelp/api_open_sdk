package schoolTime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
)

func (*Semester) GormDataType() string {
	return "string"
}

func CastSemester(in interface{}) *Semester {
	switch in.(type) {
	case int, uint, int8, int16, int32, int64:
		return &Semester{
			Num: in.(int32),
		}
	case string:
		if v, err := strconv.Atoi(in.(string)); err == nil {
			return &Semester{
				Num: int32(v),
			}
		}
	}
	return &Semester{
		Num: int32(0),
	}
}

func (x *Semester) UnmarshalJSON(data []byte) error {
	y := new(interface{})
	if err := json.Unmarshal(data, y); err != nil {
		return err
	}
	castSemester := CastSemester(y)
	if castSemester.Num == 0 {
		return errors.New("cast semester error")
	} else {
		*x = Semester{Num: castSemester.Num}
		return nil
	}
}

func (x *Semester) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Num)
}

func (x *Semester) Scan(src interface{}) error {
	castSemester := CastSemester(src)
	if castSemester.Num == 0 {
		return errors.New("cast schoolYear error")
	} else {
		*x = Semester{
			Num: castSemester.Num,
		}
		return nil
	}
}

func (x *Semester) Value() (driver.Value, error) {
	return x.Num, nil
}
