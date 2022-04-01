package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/godror/godror"
	"math"
	"strconv"
)

type BitInt struct {
	OriginInt int64
	Array     []int32
}

func (BitInt) GormDataType() string {
	return "int"
}

func (i *BitInt) Scan(src interface{}) error {
	bitIntStr, ok := src.(godror.Number)
	if !ok {
		return errors.New("int bit cannot paste in string")
	}
	bitInt, err := strconv.ParseInt(bitIntStr.String(), 10, 64)
	if err != nil {
		return err
	}
	out := BitToList(bitInt * 2)
	*i = BitInt{
		OriginInt: bitInt,
		Array:     out,
	}
	return nil
}

func (i *BitInt) Value() (driver.Value, error) {
	return i.OriginInt, nil
}

func (i *BitInt) UnmarshalJSON(bytes []byte) error {
	var bit int64
	if err := json.Unmarshal(bytes, &bit); err != nil {
		return err
	}
	out := BitToList(bit)
	*i = BitInt{
		OriginInt: bit,
		Array:     out,
	}
	return nil
}

func BitToList(bits int64) []int32 {
	out := make([]int32, 0)
	// 通过二进制计算出明细
	if bits == 0 {
		return out
	}
	t := math.Round(math.Log2(float64(bits)))
	for i := int32(0); float64(i) <= t; i++ {
		if uint(math.Pow(2, float64(i)))&uint(bits) != 0 {
			out = append(out, i)
		}
	}
	return out
}
