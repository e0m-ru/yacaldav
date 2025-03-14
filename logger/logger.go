package logger

import (
	"log"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Логгер
type Logger struct {
	level  LogLevel
	Logger *log.Logger
}

// Конструктор логгера
func NewLogger(level LogLevel, output string) *Logger {
	var writer *os.File
	var err error

	if strings.ToLower(output) != "stdout" {
		writer, err = os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
	} else {
		writer = os.Stdout
	}

	return &Logger{
		level:  level,
		Logger: log.New(writer, "", log.LstdFlags),
	}
}

// Метод для логирования информации
func (l *Logger) Info(msg string) {
	if l.level <= INFO {
		l.Logger.Printf("INFO: %s\n", msg)
	}
}

// Метод для логирования предупреждений
func (l *Logger) Warn(msg string) {
	if l.level <= WARN {
		l.Logger.Printf("WARN: %s\n", msg)
	}
}

// Метод для логирования ошибок
func (l *Logger) Error(err error) {
	if err != nil {
		l.Logger.Printf("ERROR: %s\n", err)
		panic(err)
	}
}

func (l *Logger) Print(msg string) {
	if l.level <= INFO {
		l.Logger.Printf("%s\n", msg)
	}
}
func (l *Logger) Printf(tmp string, v ...any) {
	if l.level <= INFO {
		l.Logger.Printf(tmp, v...)
	}
}
