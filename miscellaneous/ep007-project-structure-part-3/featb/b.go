package featb

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Do() error {
	fmt.Println("featb.Do() was called")
	rand.Seed(time.Now().Unix())
	if rand.Intn(100)%2 == 0 {
		return errors.New("featb.Do() error")
	}

	return nil
}
