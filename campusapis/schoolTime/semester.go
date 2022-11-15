package schoolTime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func (*Semester) GormDataType() string {
	return "string"
}

func CastSemester(in interface{}) *Semester {
	switch v := in.(type) {
	case int, uint, int8, int16, int32, int64:
		return &Semester{
			Num: v.(int32),
		}
	case float32, float64:
		return &Semester{
			Num: int32(v.(float64)),
		}
	case string:
		if v, err := strconv.Atoi(v); err == nil {
			return &Semester{
				Num: int32(v),
			}
		}
	}
	return &Semester{
		Num: int32(0),
	}
}

func (x *Semester) IsZero() bool {
	return x == nil || x.Num == 0
}

func (x *Semester) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	castSemester := CastSemester(v)
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

func (x *Semester) ShortName() string {
	if x == nil {
		return ""
	}
	return strconv.Itoa(int(x.Num))
}

func (x *Semester) FullName() string {
	return fmt.Sprintf("第%s学期", map[int32]string{
		1: "一",
		2: "二",
		3: "三",
	}[x.Num])
}
