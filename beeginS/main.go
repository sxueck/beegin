package main

import (
	"github.com/sxueck/beegin"
	"log"
	"net/http"
)

func main() {
	r := beegin.New()
	r.GET("/hello", func(c *beegin.Context) error {
		return c.String(http.StatusOK,"hello World")
	})

	r.GET("/args", func(c *beegin.Context) error {
		return c.String(http.StatusOK,c.Query("args"))
	})

	log.Fatal(r.Run(":80"))
}
