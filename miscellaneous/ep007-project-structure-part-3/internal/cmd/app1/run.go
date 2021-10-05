package app1

import (
	"fmt"

	"github.com/striversity/gotr/misc-ep007/feata"
	"github.com/striversity/gotr/misc-ep007/featb"
)

func Run() error {
	if err := feata.Do(); err != nil {
		return fmt.Errorf("failed  to initialize feata: %w", err)
	}

	if err := featb.Do(); err != nil {
		return fmt.Errorf("failed  to initialize featb: %w", err)
	}

	return nil
}
