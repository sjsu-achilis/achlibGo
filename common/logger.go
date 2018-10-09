package common

import (
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	l = NewLogger()
}

//Logger ...
type Logger struct {
	outf *os.File
	ent  *logrus.Entry
}

//NewLogger ...
func NewLogger() *Logger {
	l := new(Logger)
	l.ent = logrus.NewEntry(logrus.New())
	l.ent.Logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, ForceColors: false})
	//l.ent.Logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	l.ent.Logger.SetLevel(logrus.DebugLevel)
	return l
}

//SetLogOutputFile ...
func SetLogOutputFile(fname string) { l.SetLogOutputFile(fname) }

//SetLogOutputFile ...
func (l *Logger) SetLogOutputFile(fname string) {
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
func Log(m ...map[string]interface{}) *logrus.Entry { return l.Log(m...) }

//Log ...
func (l *Logger) Log(m ...map[string]interface{}) *logrus.Entry {
	data := make(map[string]interface{})
	if m != nil {
		data = m[0]
	}
	if pc, file, line, ok := runtime.Caller(2); ok {
		fName := runtime.FuncForPC(pc).Name()
		data["file"], data["line"], data["func"] = file, line, fName
	}
	l.ent.Data = data

	return l.ent
}
