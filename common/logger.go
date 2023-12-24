package common

import (
	"log"
)

func SilentError(v ...any) {
	log.Println(v...)
}
