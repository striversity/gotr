package foo

import "fmt"

// Println is my simple println function
func Println(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Print(x.(bool))
	case int:
		fmt.Print(x.(int))
	case float64:
		fmt.Print(x.(float64))
	case complex128:
		fmt.Print(x.(complex128))
	case string:
		fmt.Print(x.(string))
	// case Person:
	// 	fmt.Print(x.(Person))
	case chan int:
		fmt.Print(x.(chan int))
	default:
		fmt.Print("Unknown type")
	}

	fmt.Print("\n")
}
