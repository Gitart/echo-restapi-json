package main

import (
	"api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type magazines []models.Magazine

func readJson() magazines {
	file, err := ioutil.ReadFile("./document.json")
	if err != nil {
		log.Fatal(err)
	}
	m := magazines{}
	err = json.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func writeJson(data magazines) {
	json_to_file, _ := json.Marshal(data)
	err := ioutil.WriteFile("./document.json", json_to_file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()

	e.GET("/magazines", func(c echo.Context) error {
		m := readJson()

		return c.JSON(http.StatusOK, m)
	})

	e.GET("/magazines/:id", func(c echo.Context) error {
		m := readJson()

		for _, magazine := range m {
			if c.Param("id") == strconv.Itoa(magazine.Id) {
				return c.JSON(http.StatusOK, magazine)
			}
		}
		return c.String(http.StatusBadRequest, "Bad request.")
	})

	e.POST("/magazines", func(c echo.Context) error {
		m := readJson()

		new_magazine := new(models.Magazine)
		err := c.Bind(new_magazine)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request.")
		}

		m = append(m, *new_magazine)

		writeJson(m)
		return c.JSON(http.StatusOK, m)
	})

	e.PUT("/magazines/:id", func(c echo.Context) error {
		m := readJson()

		updated_magazine := new(models.Magazine)
		err := c.Bind(updated_magazine)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request.")
		}
		for i, magazine := range m {
			if strconv.Itoa(magazine.Id) == c.Param("id") {
				m = append(m[:i], m[i+1:]...)
				m = append(m, *updated_magazine)

				writeJson(m)

				return c.JSON(http.StatusOK, m)
			}
		}

		return c.String(http.StatusBadRequest, "Bad request.")
	})

	e.DELETE("/magazines/:id", func(c echo.Context) error {
		m := readJson()

		for i, magazine := range m {
			if strconv.Itoa(magazine.Id) == c.Param("id") {
				m = append(m[:i], m[i+1:]...)
				writeJson(m)

				return c.JSON(http.StatusOK, m)
			}
		}
		return c.String(http.StatusBadRequest, "Bad request.")
	})

	e.Start(":5000")
}
