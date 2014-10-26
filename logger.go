package log

import (
	_l "log"
)

type App struct{}
type Error struct{}
type Profiling struct{}

type Logger interface {
	Printf(pattern string, params ...interface{})
	Println(msg string, params ...interface{})
}

func Printf(pattern string, params ...interface{}) {
	_l.Printf(pattern, params)
}
func Println(msg string, params ...interface{}) {
	_l.Println(msg, params)
}
