package tools

import (
	"domino-api-gin/errors"
	"fmt"
	"reflect"
	"strconv"
)

// StringToInt 文字转数字
func StringToInt(v string) (intValue int) {
	intValue, err := strconv.Atoi(v)
	if err != nil {
		panic(errors.ThrowForbidden("id必须是数字"))
	}

	return
}

// StringToUint 文字转无符号数字
func StringToUint(v string) (uintValue uint) {
	intValue := StringToInt(v)
	uintValue = uint(intValue)
	return
}

// ThrowErrorWhenIsEmpty 判断是否为空
func ThrowErrorWhenIsEmpty(ins interface{}, class interface{}, name string) (isEmpty bool) {
	isEmpty = reflect.DeepEqual(ins, class)

	if name != "" {
		if isEmpty {
			panic(errors.ThrowEmpty(fmt.Sprintf("%v不存在", name)))
		}
	}

	return isEmpty
}

// ThrowErrorWhenIsRepeat 判断是否重复
func ThrowErrorWhenIsRepeat(ins interface{}, class interface{}, name string) (isRepeat bool) {
	isRepeat = !reflect.DeepEqual(ins, class)

	if name != "" {
		if isRepeat {
			panic(errors.ThrowForbidden(fmt.Sprintf("%v重复", name)))
		}
	}

	return
}
