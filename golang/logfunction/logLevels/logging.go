package main

import (
	"os"
	//logrus supports log levels like Trace, Debug, Info, Warn, Error, Fatal and Panic
	log "github.com/sirupsen/logrus"

)

func main() {
	// If the file doesn't exist,this will create the file 
	//or if the file exists this will overwrite to the file
	file, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	//Debug level message will onl be printed if log.level is equal to log.debuglevel
	log.SetLevel(log.DebugLevel)
	log.Debug("This is debug level: Hello everyone.")
	log.Info("This is info level: Welcome to roost.ai!")
	log.Warn("This is warning level: Complete the assignment before tuesday")
	log.Error("This is error level: I dont know much about golang! but I'm not quitting.")
}
