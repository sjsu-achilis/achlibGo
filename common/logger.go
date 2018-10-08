package common

import (
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

//Logger ...
type Logger struct {
	outf *os.File
	ent  *logrus.Entry
}

//NewLogger ...
func NewLogger() *Logger {
	l := new(Logger)
	l.ent = logrus.NewEntry(logrus.New())
	l.ent.Logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	l.ent.Logger.SetLevel(logrus.DebugLevel)
	return l
}

//SetOutputFile ...
func (l *Logger) SetOutputFile(fname string) {
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		l.ent.Error(err)
		os.Exit(1)
	}
	w := io.MultiWriter(f, os.Stdout)
	l.ent.Logger.SetOutput(w)
	l.outf = f
}

//Log ...
func (l *Logger) Log(m ...map[string]interface{}) *logrus.Entry {
	if len(m) != 0 {
		for k, v := range m[0] {
			l.ent = l.ent.WithField(k, v)
		}
	}
	if pc, file, line, ok := runtime.Caller(1); ok {
		fName := runtime.FuncForPC(pc).Name()
		return l.ent.WithField("file", file).WithField("line", line).WithField("func", fName)
	}

	return l.ent
}
