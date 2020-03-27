package main

import "benc/model"
import "fmt"

func main() {
	cr1 := model.SearchReq{
		Q:      "books",
		Params: []string{"offset=10", "limit=100", "pubyear=2019"},
	}

	fmt.Printf("cr1: %v\n", cr1)
}
