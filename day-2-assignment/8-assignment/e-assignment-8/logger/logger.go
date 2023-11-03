package logger

import "log"

type LoggerStore struct { // this is a loggers struct
	log *log.Logger
}

func New() LoggerStore { // this i snew fuc which return loggerstore
	return LoggerStore{
		log: log.New(log.Writer(), "siddarth", 56),
	}
}
