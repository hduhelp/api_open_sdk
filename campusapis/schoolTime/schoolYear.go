package schoolTime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (*SchoolYear) GormDataType() string {
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

func (x *SchoolYear) IsZero() bool {
	return x == nil || x.Year == 0
}

func (x *SchoolYear) ShortName() string {
	return strconv.Itoa(int(x.Year))
}

func (x *SchoolYear) FullName() string {
	return fmt.Sprintf("%d-%d", x.Year, x.Year+1)
}

func (x *SchoolYear) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	st := CastSchoolYear(v)
	if st.Year == 0 {
		return errors.New("cast schoolYear error")
	} else {
		*x = SchoolYear{
			Year: st.Year,
		}
		return nil
	}
}

func (x *SchoolYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.FullName())
}

func (x *SchoolYear) Scan(src interface{}) error {
	st := CastSchoolYear(src)
	if st.Year == 0 {
		return errors.New("cast schoolYear error")
	} else {
		*x = SchoolYear{
			Year: st.Year,
		}
		return nil
	}
}

func (x *SchoolYear) Value() (driver.Value, error) {
	return x.FullName(), nil
}
