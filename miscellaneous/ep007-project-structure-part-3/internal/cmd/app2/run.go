package app2

import (
	"fmt"

	"github.com/striversity/gotr/misc-ep007/feata"
	"github.com/striversity/gotr/misc-ep007/featc"
)

func Run() error {
	if err := feata.Do(); err != nil {
		return fmt.Errorf("failed  to initialize feata: %w", err)
	}

	if err := featc.Do(); err != nil {
		return fmt.Errorf("failed  to initialize featc: %w", err)
	}

	return nil
}
