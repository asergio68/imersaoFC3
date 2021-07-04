package main

import (
	"github.com/asergio68/imersaoFC3/goapp/data"
	webserver2 "github.com/asergio68/imersaoFC3/goapp/webserver"
)

func main_() {
	data.LoadData()
	webserver := webserver2.NewWebServer()
	webserver.Serve()
}