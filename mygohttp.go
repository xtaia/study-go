package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handlerindex)
	log.Println("starting echo")
	err := e.Start(":8080")
	if err != nil {
		log.Fatal("echo", err)
	}
}

func handlerindex(c echo.Context) error {
	log.Println("hello world handlerindex")
	return c.JSON(http.StatusOK, `{"hello":"world"}`)
}
