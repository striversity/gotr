package main

import (
	"fmt"

	"github.com/striversity/misc007/app1/internal/stuff/secret1"
	"github.com/striversity/misc007/app1/pkga"
)
import "github.com/striversity/misc007/app1/some/other/leaker/leaker2"

func main() {
	fmt.Println(pkga.PUBLIC)
	fmt.Println(secret1.SECRET)
	fmt.Println(leaker2.GetSecret())
}
