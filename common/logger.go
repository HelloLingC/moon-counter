package common

import (
	"log"
)

type Logger struct {
	path string
}

func (l Logger) InitLogger(path string) {

}

func SilentError(v ...any) {
	log.Println(v...)
}

func Info(v ...any) {
	log.Println(v...)
}