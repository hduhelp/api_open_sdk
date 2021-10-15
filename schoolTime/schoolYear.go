package schoolTime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (SchoolYear) GormDataType() string {
	return "string"
}

func CastSchoolYear(in interface{}) *SchoolYear {
	switch in.(type) {
	case int, uint, int8, int16, int32, int64:
		return &SchoolYear{
			Year: in.(int32),
		}
	case string:
		if arr := strings.Split(in.(string), "-"); len(arr) != 0 {
			if y, err := strconv.Atoi(arr[0]); err == nil {
				return &SchoolYear{
					Year: int32(y),
				}
			}
		}
	}
	return &SchoolYear{
		Year: int32(0),
	}
}

func (m SchoolYear) ShortName() string {
	return strconv.Itoa(int(m.Year))
}

func (m SchoolYear) FullName() string {
	return fmt.Sprintf("%d-%d", m.Year, m.Year+1)
}

func (m *SchoolYear) UnmarshalJSON(data []byte) error {
	y := new(interface{})
	if err := json.Unmarshal(data, y); err != nil {
		return err
	}
	st := CastSchoolYear(y)
	if st.Year == 0 {
		return errors.New("cast schoolYear error")
	} else {
		*m = *st
		return nil
	}
}

func (m SchoolYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.FullName())
}

func (m *SchoolYear) Scan(src interface{}) error {
	st := CastSchoolYear(src)
	if st.Year == 0 {
		return errors.New("cast schoolYear error")
	} else {
		*m = *st
		return nil
	}
}

func (m SchoolYear) Value() (driver.Value, error) {
	return m.FullName(), nil
}
