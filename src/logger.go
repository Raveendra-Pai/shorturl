package main

import (
	"log"
	"os"
	"strings"
)

type ILogger interface {
	Info(msg string)
	Error(msg string)
}

type FileLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *FileLogger) Init(filename string) error {

	//Only in windows application binary has exe extension
	isWindows := strings.HasSuffix(os.Args[0], ".exe")

	var tmpPath = "/app/logs/"

	if isWindows == true {
		tmpPath = "./../logs/"
	}

	fullfilename := tmpPath + filename
	file, err := os.OpenFile(fullfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	l.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func (l FileLogger) Info(msg string) {
	l.infoLogger.Println(msg)
}

func (l FileLogger) Error(msg string) {
	l.errorLogger.Println(msg)
}
