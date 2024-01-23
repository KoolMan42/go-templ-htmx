package handler

import (
	"github.com/koolman42/go-templ-htmx/view/index"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct {}


func(h IndexHandler) HandleIndexShow(c echo.Context) error {
  return render(c,index.Show("Thomas"))
}
