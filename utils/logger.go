package utils

import (
	"bytes"
	"log"

	"github.com/fatih/color"
)

// DBlogger Database err logger
func DBlogger() *log.Logger {
	var buf bytes.Buffer
	lgr := log.New(&buf, color.RedString("[DBERR]"), log.Lshortfile)
	return lgr
}
