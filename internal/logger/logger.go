package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	log *log.Logger
}

func New() *Logger {
	return &Logger{
		log.Default(),
	}
}

func (log *Logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	log.log.Printf("[Error]: %s\n", msg)

}

func (log *Logger) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	log.log.Printf("[Info]: %s\n", msg)
}
