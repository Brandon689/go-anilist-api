package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rl404/verniy"
	"log"
	"net/http"
)

func searchAnime(c echo.Context) error {

	search := c.QueryParam("search")

	data, err := v.SearchAnime(verniy.PageParamMedia{
		Search: search,
	}, 1, 20)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, data)
}
