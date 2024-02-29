package filelogger

import (
	"fmt"
	"time"
)

func CreateLogFileName() string {
	return fmt.Sprintf("%s.log", time.Now().Format("2006-01-02:15:04:05"))
}
