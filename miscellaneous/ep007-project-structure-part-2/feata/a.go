package feata

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Do() error {
	fmt.Println("feata.Do() was called")
	rand.Seed(time.Now().Unix())
	if rand.Intn(100)%2 == 0 {
		return errors.New("feata.Do() error")
	}

	return nil
}
