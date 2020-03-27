package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	dumpData(PeopleData)

	s1 := dataToStruct(PeopleData[0])
	fmt.Printf("%#v\n", s1)
}

func dataToStruct(data []TableField) (strt interface{}) {
	// number of fields required
	fields := make([]reflect.StructField, len(data))
	var t reflect.Type

	for i, tf := range data {
		fields[i].Name = tf.Name
		switch tf.Type {
		case "int":
			t = reflect.TypeOf(int(0))
		case "uint8":
			t = reflect.TypeOf(uint(0))
		case "string":
			t = reflect.TypeOf(string(""))
		case "float64":
			t = reflect.TypeOf(float64(0.0))
		}
		fields[i].Type = t
	}

	typ := reflect.StructOf(fields)
	v := reflect.New(typ).Elem()

	for i, tf := range data {
		switch tf.Type {
		case "int":
			tv, _ := strconv.ParseInt(tf.Value, 10, 64)
			v.Field(i).SetInt(tv)
		case "uint8":
			tv, _ := strconv.ParseUint(tf.Value, 10, 64)
			v.Field(i).SetUint(tv)
		case "string":
			v.Field(i).SetString(tf.Value)
		case "float64":
			tv, _ := strconv.ParseFloat(tf.Value, 64)
			v.Field(i).SetFloat(tv)
		}
	}

	s := v.Addr().Interface()
	return s
}
