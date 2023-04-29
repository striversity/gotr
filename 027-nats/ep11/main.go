package main

import "fmt"

func main() {
	// key (Track #) -> value (Title)
	// "001"  -> "Three Little Birds"
	// "002"  -> "The Wind Cries Mary
	// "003"  -> "The Birds"

	playlist := make(map[string]string)

	playlist["001"] = "Three Little Birds"

	playlist["002"] = "The Wind Cries Mary"

	playlist["003"] = "The Birds"

	for tr, title := range playlist {
		fmt.Printf("Track #%v -> Title: %v\n", tr, title)
	}
}
