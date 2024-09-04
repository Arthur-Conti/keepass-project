package config

import (
	"github.com/fatih/color"
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
}

func newLogger() *Logger {
	writter := io.Writer(os.Stdout)
	greenLog := color.New(color.FgGreen).SprintFunc()
	yellowLog := color.New(color.FgYellow).SprintFunc()
	redLog := color.New(color.FgRed).SprintFunc()
	return &Logger{
		debug: log.New(writter, greenLog("DEBUG: "), log.Ldate|log.Ltime),
		info: log.New(writter, greenLog("INFO: "), log.Ldate|log.Ltime),
		warn: log.New(writter, yellowLog("WARN: "), log.Ldate|log.Ltime),
		err: log.New(writter, redLog("ERROR: "), log.Ldate|log.Ltime),
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}


