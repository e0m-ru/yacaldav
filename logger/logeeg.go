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
	logger *log.Logger
}

// Конструктор логгера
func NewLogger(level LogLevel, output string) (*Logger, error) {
	var writer *os.File
	var err error

	if strings.ToLower(output) == "file" {
		writer, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
	} else {
		writer = os.Stdout
	}

	return &Logger{
		level:  level,
		logger: log.New(writer, "", log.LstdFlags),
	}, nil
}

// Метод для логирования информации
func (l *Logger) Info(msg string) {
	if l.level <= INFO {
		l.logger.Printf("INFO: %s\n", msg)
	}
}

// Метод для логирования предупреждений
func (l *Logger) Warn(msg string) {
	if l.level <= WARN {
		l.logger.Printf("WARN: %s\n", msg)
	}
}

// Метод для логирования ошибок
func (l *Logger) Error(err error) {
	if err != nil {
		// l.logger.Printf("ERROR: %s\n", err)
		panic(err)
	}
}
