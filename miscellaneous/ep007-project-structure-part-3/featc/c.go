package featc

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Do() error {
	fmt.Println("featc.Do() was called")
	rand.Seed(time.Now().Unix())
	if rand.Intn(100)%2 == 0 {
		return errors.New("featc.Do() error")
	}

	return nil
}
