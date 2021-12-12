package leaker2

import "github.com/striversity/misc007/app1/some/other/internal/secret2"

func GetSecret() string {
	return secret2.SECRET
}
