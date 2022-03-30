package campusapis

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/grpc/codes"
)

func Require(req any, list ...any) error {
	emptyList := make([]string, 0)
	for _, v := range list {
		if reflect.ValueOf(v).IsZero() {
			emptyList = append(emptyList, GetFieldName(req, v))
		}
	}
	if len(emptyList) != 0 {
		return RequiredButEmpty{values: emptyList}
	}
	return nil
}

func GetFieldName(structPoint any, fieldPinter any) (name string) {
	valStruct := reflect.ValueOf(structPoint).Elem()
	valField := reflect.ValueOf(fieldPinter).Elem()

	for i := 0; i < valStruct.NumField(); i++ {
		valueField := valStruct.Field(i)
		if valueField.Addr().Interface() == valField.Addr().Interface() {
			return valStruct.Type().Field(i).Name
		}
	}
	return
}

func RequireWarped(req any, list ...any) error {
	if err := Require(req, list[1:]...); err != nil {
		return Status_INVALID_ARGUMENT.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}

type RequiredButEmpty struct {
	values []string
}

func (e RequiredButEmpty) Error() string {
	return fmt.Sprintf("%s are required but empty", strings.Join(e.values, ", "))
}
