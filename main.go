package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tokokurma/config"
	"tokokurma/route"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("run")
	go func() {
		err := route.Run(gin.Default(), ":"+config.PORT)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("serving at port", config.PORT)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}
