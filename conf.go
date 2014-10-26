package log

import ()

type conf struct {
	App       logFile
	Error     logFile
	Profiling logFile
}

type logFile struct {
	Name string
	Path string
}
