package common

import (
	"log"
)

func HasError(err error) bool {
	if err != nil {
		log.Printf("Encountered error: %v", err)
		return true
	}
	return false
}
