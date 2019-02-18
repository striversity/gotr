package foo

import (
	"fmt"
	"reflect"
	"strings"
)

// Println is my simple println function
func Println(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	switch t.Kind() {
	case reflect.Struct:
		s := printStructExpanded(x)
		fmt.Print(s)
	case reflect.Ptr:
		v2 := reflect.Indirect(v)
		s := printStructExpanded(v2.Interface())
		fmt.Printf("&%v", s)
	default:
		fmt.Printf("Unknown type")
	}

	fmt.Print("\n")
}

func printStructExpanded(x interface{}) string {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t.Kind() != reflect.Struct {
		return "<nil>"
	}

	sb := &strings.Builder{}
	n := t.NumField()
	fmt.Fprintf(sb, "%v{", t)
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)
		fmt.Fprintf(sb, "%v: %v, ", tt.Name, vv)
	}
	fmt.Fprintln(sb, "}")

	return sb.String()
}
