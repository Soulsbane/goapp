package filelogger

import (
	"log"
	"os"
	"path"
)

type FileLogger struct {
	logger *log.Logger
}

func New(fileName string, directory string, flag int) FileLogger {
	os.MkdirAll(directory, os.ModePerm)
	file, err := os.OpenFile(path.Join(directory, fileName), flag|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	var fileLogger FileLogger
	fileLogger.logger = log.New(file, "", log.LstdFlags)

	return fileLogger
}

func (fileLogger *FileLogger) Println(msg ...interface{}) {
	fileLogger.logger.Println(msg...)
}

func (fileLogger *FileLogger) Print(msg ...interface{}) {
	fileLogger.logger.Print(msg...)
}
