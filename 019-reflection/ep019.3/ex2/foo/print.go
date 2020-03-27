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
	case reflect.Struct:
		printStructExpanded(x)
	case reflect.Ptr:
		v2 := reflect.Indirect(v)
		printStructExpanded(v2.Interface())
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
		vv := v.Field(i)
		fmt.Printf("%v: %v, ", tt.Name, vv)
	}
	fmt.Println("}")
}
