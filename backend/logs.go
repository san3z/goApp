package backend

import (
	"log"
	"os"
)

var logger *log.Logger

var logFile string = "backend/logs/promo.log"

func CritSize() error {
	file, err := os.OpenFile(logFile, os.O_RDWR, 0o644)
	defer file.Close()
	if err != nil {
		return err
	}

	file.Truncate(0)
	file.Close()
	os.Remove(logFile)
	return nil
}

func initLogger() {
	//	fmt.Println("initLogger kicks in")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %d", logFile, err)
		return
	}
	fileInfo, _ := os.Stat(logFile)
	fileSize := fileInfo.Size()
	logger = log.New(file, "PROMO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if fileSize > 10000 {
		CritSize()
	}
}
