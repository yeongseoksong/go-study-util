package util

import (
	"fmt"
	"reflect"
)

func IsEqual(x any, y any) bool {
	flag := x == y
	Print(flag)
	return flag
}

func PrintAll(params ...any) {
	for _, p := range params {
		Print(p)
	}
}

func PrintAllWithFlag(flag bool, params ...any) {
	for _, p := range params {
		PrintWithFlag(flag, p)
	}
}

func Print(v any) {
	PrintWithFlag(false, v)
}

func PrintWithFlag(flag bool, v any) {
	PrintUnderline()
	fmt.Printf("** %v 정보 출력 : ** \n", v)
	PrintUnderline()

	r := reflect.ValueOf(v)

	switch r.Kind() {
	case reflect.Array,
		reflect.Slice:
		fmt.Printf("%v len:%d cap:%d\n",
			r.Interface(),
			r.Len(),
			r.Cap(),
		)

	default:
		fmt.Println(v)
	}

	if flag {
		IsNil(v)
	}
	PrintUnderline()

}

func IsNil(x any) {
	fmt.Printf("%v isNil : %v \n", x, isNil(x))
}

func PrintUnderline() {
	fmt.Println("----------------------------------------")
}
