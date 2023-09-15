package main

import (
	"log"
	"os"
	"time"

	info_controller "github.com/thaessaep/rpc/info-controller"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(file, "INFO\t", log.Ldate|log.Ltime)

	infoController := info_controller.New(logger)

	timer := time.NewTimer(5 * time.Minute)
	for {
		select {
		case <-timer.C:
			infoController.Last5MinutesNews()
		}
	}
}
