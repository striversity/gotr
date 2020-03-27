package foo

import (
	"fmt"
	"reflect"
)

// Println is my simple println function
func Println(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	switch t.Kind() {
	case reflect.Bool:
		fmt.Printf("boolean value '%v'\n", v)
	case reflect.Int:
		fmt.Printf("int value '%v'\n", v)
	case reflect.Float64:
		fmt.Printf("float64 value '%v'\n", v)
	case reflect.Complex128:
		fmt.Printf("complex128 value '%v'\n", v)
	case reflect.String:
		fmt.Printf("string value '%v'\n", v)
	case reflect.Struct:
		printStructExpanded(x)
	case reflect.Chan:
		fmt.Printf(v.String())
	default:
		fmt.Printf("Unknown type")
	}

	fmt.Print("\n")
}

func printStructExpanded(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t.Kind() != reflect.Struct {
		fmt.Printf("ERR: Not a struct, expected struct value of 'kind' struct...\n")
		return
	}

	n := t.NumField()
	fmt.Printf("'%v'{", t)
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		vv:= v.Field(i)
		fmt.Printf("%v: %v, ", tt.Name, vv)
	}
	fmt.Println("}")
}
