package glog

import (
	"io"
	"log"
	"os"
)

type DefaultLogger struct {
	*log.Logger
}

func newDefaultLogger(out io.Writer, prefix string, flag int) *DefaultLogger {
	s := &DefaultLogger{}
	s.Logger = log.New(os.Stderr, "", log.LstdFlags)

	return s
}

func (this *DefaultLogger) Error(args ...interface{}) {
	this.Print(args)
}

func (this *DefaultLogger) Errorf(format string, args ...interface{}) {
	this.Printf(format, args)
}

func (this *DefaultLogger) Warn(args ...interface{}) {
	this.Print(args)
}

func (this *DefaultLogger) Warnf(format string, args ...interface{}) {
	this.Printf(format, args)
}

func (this *DefaultLogger) Info(args ...interface{}) {
	this.Print(args)
}

func (this *DefaultLogger) Infof(format string, args ...interface{}) {
	this.Printf(format, args)
}

func (this *DefaultLogger) Debug(args ...interface{}) {
	this.Print(args)
}

func (this *DefaultLogger) Debugf(format string, args ...interface{}) {
	this.Printf(format, args)
}
