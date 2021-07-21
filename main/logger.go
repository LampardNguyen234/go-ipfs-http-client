package main

import (
	"fmt"
	"log"
	"os"
)

// ServerLogger implements a logger for the server.
type ServerLogger struct {
	Log      *log.Logger
	IsEnable bool
}

// Printf overrides log.Printf.
func (l ServerLogger) Printf(format string, v ...interface{}) {
	if l.IsEnable {
		l.Log.Printf(format, v...)
	}
}

// Println overrides log.Println.
func (l ServerLogger) Println(v ...interface{}) {
	if l.IsEnable {
		l.Log.Println(v...)
	}
}

// Fatalf overrides log.Fatalf.
func (l ServerLogger) Fatalf(format string, v ...interface{}) {
	if l.IsEnable {
		l.Log.Fatalf(format, v...)
	}
}

// Fatalln overrides log.Fatalln.
func (l ServerLogger) Fatalln(v ...interface{}) {
	if l.IsEnable {
		l.Log.Fatalln(v...)
	}
}

// Errorf prints the error with the given format.
func (l ServerLogger) Errorf(format string, v ...interface{}) {
	if l.IsEnable {
		l.Log.Printf("[ERR] " + format, v...)
	}
}

// Error prints the error with the given format.
func (l ServerLogger) Error(v ...interface{}) {
	if l.IsEnable {
		l.Log.Printf("[ERR] ")
		l.Log.Println(v...)
	}
}

// NewLogger creates a new ServerLogger. If isEnable = true, it will do logging.
// If logFile is set, it will store logging information into the given logFile.
func NewLogger(isEnable bool, logFile ...string) *ServerLogger {
	writer := os.Stdout
	if len(logFile) != 0 {
		var err error
		writer, err = os.OpenFile(logFile[0], os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
	}
	Log := log.New(writer, "", log.Ldate|log.Ltime)

	return &ServerLogger{
		Log:      Log,
		IsEnable: isEnable,
	}
}

// Logger is the main logger of this package.
var Logger = NewLogger(true)
