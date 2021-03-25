package handlers

import "log"

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l*log.Logger) *Goodbye  {
	return &Goodbye{l}
}