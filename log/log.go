package log

import "log"

type Logger struct {
	Name string
}

func (self Logger) Log(str string) {
	log.Printf("[%v] %s", self.Name, str)
}

func InitLog(l Logger) *Logger {
	return &Logger{Name: l.Name}
}
