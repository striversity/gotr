package main

import (
	"fmt"
	"reflect"
)

func main() {
	dumpData(PeopleData)

	s1 := dataToStruct(PeopleData[0])
	fmt.Printf("%#v\n", s1)
}

func dataToStruct(data []TableField) interface{} {
	var structFields []reflect.StructField
	// number of struct fields required
	structFields = make([]reflect.StructField, len(data))

	for i, tf := range data {
		structFields[i].Name = tf.Name
	}

	typ := reflect.StructOf(structFields)
	v := reflect.New(typ).Elem()
	s := v.Addr().Interface()
	return s
}
