package main

import (
	"fmt"
)

type argError struct {
	arg  string
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("[FATAL] %s: %s", e.arg, e.prob)
}
