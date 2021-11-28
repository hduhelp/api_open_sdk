package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

func DateFromString(str string) (*Date, error) {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return nil, err
	}
	return &Date{
		Time:   t,
		string: str,
	}, nil
}

func DateFromTime(t time.Time) *Date {
	return &Date{
		Time: t,
	}
}

type Date struct {
	time.Time
	string
}

func (Date) GormDataType() string {
	return "string"
}

func (d Date) String() string {
	if d.string != "" {
		return d.string
	}
	return d.Format("2006-01-02")
}
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Date) UnmarshalJSON(bytes []byte) error {
	var str string
	err := json.Unmarshal(bytes, &str)
	if err != nil {
		return err
	}
	nd, err := DateFromString(str)
	if err != nil {
		return err
	}
	*d = *nd
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d *Date) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("cannot parse src into string")
	}
	nd, err := DateFromString(str)
	if err != nil {
		return err
	}
	*d = *nd
	return nil
}
