package util

import "reflect"

func isNil(x any) bool {
	if x == nil {
		return true
	}

	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Ptr,
		reflect.Map,
		reflect.Slice,
		reflect.Chan,
		reflect.Func,
		reflect.Interface:
		return v.IsNil()
	}
	return false

}
