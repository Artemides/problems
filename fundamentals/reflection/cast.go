package reflection

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type Temp float64

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return value.Type().String() + " 0x" + strconv.FormatUint(uint64(value.Pointer()), 16)
	default:
		return value.Type().String() + " value"
	}
}

func ReflecMain() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	var _var Temp
	s := reflect.TypeOf(_var)
	fmt.Printf("%s", s.String())
}
