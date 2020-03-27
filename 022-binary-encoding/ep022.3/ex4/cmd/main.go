package main

import "benc/model"
import "fmt"

func main() {
	cr1 := model.SearchReq{
		Q:      "books",
		Params: map[string]string{"offset": "10", "limit": "100", "pubyear": "2019"},
	}

	sr1 := model.SearchResp{
		Id:     5,
		Ans:    "Book 1: Book 2",
		Source: model.From_Current,
	}

	fmt.Printf("cr1: %v\n", cr1)
	fmt.Printf("sr1: %v\n", sr1)
}
