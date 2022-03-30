package campusapis

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
)

// Require 判断入参是否为空值
func Require(req any, list ...string) error {
	emptyList := make([]string, 0)
	val := reflect.ValueOf(req).Elem()
	if !val.IsValid() {
		return nil
	}
	lowerKeysMap := lo.KeyBy(list, func(str string) string { return strings.ToLower(str) })
	for _, index := range lo.Range(val.NumField()) {
		fieldName := val.Type().Field(index).Name
		if key, exists := lowerKeysMap[strings.ToLower(fieldName)]; exists {
			emptyList = append(emptyList, key)
		}
	}
	if len(emptyList) != 0 {
		return RequiredButEmpty{values: emptyList}
	}
	return nil
}

func RequireWarped(req any, list ...string) error {
	if err := Require(req, list...); err != nil {
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
