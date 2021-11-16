package main

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type magazines []models.Magazine

func main() {
	e := echo.New()

	lot := magazines{
		{
			Id:      1,
			Title:   "Top models",
			Company: "Brezz",
			Price:   29.99,
			Month:   12,
			Year:    2020,
		},
		{
			Id:      2,
			Title:   "World ranking",
			Company: "Vuzz",
			Price:   19.99,
			Month:   05,
			Year:    2019,
		},
	}

	e.GET("/magazines", func(c echo.Context) error {
		return c.JSON(http.StatusOK, lot)
	})

	e.GET("/magazines/:id", func(c echo.Context) error {
		for _, magazine := range lot {
			if c.Param("id") == strconv.Itoa(magazine.Id) {
				return c.JSON(http.StatusOK, magazine)
			}
		}
		return c.String(http.StatusBadRequest, "Bad request.")
	})

	e.POST("/magazines", func(c echo.Context) error { 
		new_magazine := new(models.Magazine)
		err := c.Bind(new_magazine)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request.")
		}

		lot = append(lot, *new_magazine)
		return c.JSON(http.StatusOK, lot)
	})

	e.Start(":5000")
}
