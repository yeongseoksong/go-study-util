package util

import (
	"fmt"
	"reflect"
)

func Print(x any, flags ...bool) {
	PrintUnderline()
	fmt.Printf("** %v 정보 출력 : ** \n", x)
	PrintUnderline()

	nilFlag := false
	if len(flags) > 0 {
		nilFlag = true
	}

	v := reflect.ValueOf(x)

	switch v.Kind() {
	case reflect.Array,
		reflect.Slice:
		fmt.Printf("%v len:%d cap:%d\n",
			v.Interface(),
			v.Len(),
			v.Cap(),
		)

	default:
		fmt.Println(x)
	}

	if nilFlag {
		IsNil(x)
	}
	PrintUnderline()

}

func IsNil(x any) {
	fmt.Printf("%v isNil : %v \n", x, isNil(x))
}

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

func PrintUnderline() {
	fmt.Println("----------------------------------------")
}
