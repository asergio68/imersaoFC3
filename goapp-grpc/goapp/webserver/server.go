package webserver

import (
	"net/http"

	"github.com/asergio68/imersaoFC3/goapp/data"
	"github.com/labstack/echo/v4"
)

type WebServer struct {

}

func NewWebServer() (*WebServer) {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.Logger.Fatal(e.Start(":8085"))
}

func (w WebServer) getAll(c echo.Context) error{
	return c.JSON(http.StatusOK, data.Products)
}