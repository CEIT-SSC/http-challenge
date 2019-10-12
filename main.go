package main

import "github.com/smf8/http-challange/router"

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8080"))
}
