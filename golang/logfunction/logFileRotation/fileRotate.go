package main

import (
    "fmt"
    "log"
    "time"
    rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func initiLogger() {
	//change the necessary path to your local directory
	path := "/home/chaitra/roostAsmt/golang/logfunction/logFileRotation/LogRotationFiles/old.UTC"
	writer, err := rotatelogs.New(
        fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S"),
	//change the necessary path to your local directory
        rotatelogs.WithLinkName("/home/chaitra/roostAsmt/golang/logfunction/logFileRotation/LogRotationFiles/current"),
	//MaxAge to 10 seconds means the file will be deleted after 10 seconds.
        rotatelogs.WithMaxAge(time.Second*10),
	//RotationTime is 1 second which means the file will be rotated every second
        rotatelogs.WithRotationTime(time.Second*1),
    )
    if err != nil {
        log.Fatalf("Failed to Initialize Log File %s", err)
    }
    log.SetOutput(writer)
    return
}

func main() {
    initiLogger()
    for i := 0; i < 100; i++ {
        time.Sleep(time.Second * 1)
        log.Printf("Hello, Roost")
    }
    fmt.Scanln()
}
