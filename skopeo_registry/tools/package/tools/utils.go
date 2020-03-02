package tools

import (
	"log"
	"strings"
)

func splitCredentials(s string) []string {
	slc := strings.Split(s, ":")
	if len(slc) != 2 {
		log.Fatal("Check your credentials - format should be 'user:password'")
	}
	return slc
}
