package common

var l *Logger
var c *Config
var d *Db

func init() {
	l = NewLogger()
	c = NewConfig()
	d = NewDb()
}
