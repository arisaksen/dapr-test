package main

import (
	"github.com/arisaksen/api1/author"
	"github.com/arisaksen/api1/renderer"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func GetAuthors() ([]author.Author, error) {
	authors := []author.Author{
		{Name: "J.R.R Tolkien", YearOfBirth: 1982},
		{Name: "Dan Brown", YearOfBirth: 1964},
		{Name: "Sun Tzu", YearOfBirth: -544},
	}

	return authors, nil
}

func handleGetAuthor(c echo.Context) error {
	authors, err := GetAuthors()
	if err != nil {
		return err
	}
	firstAuthor := authors[0]
	response := c.Render(http.StatusOK, "index", firstAuthor)
	return response
}

func handleGetAuthorApi(c echo.Context) error {
	user, err := GetAuthors()
	if err != nil {
		return err
	}
	contentType := c.Request().Header.Get("Content-Type")

	var response error
	switch contentType {
	case "application/xml":
		response = c.XML(http.StatusOK, user)
	case "application/json":
		fallthrough
	default:
		response = c.JSON(http.StatusOK, user)
	}

	return response
}

func main() {
	e := echo.New()
	runEnv := os.Getenv("ENVIRONMENT")
	var paths string
	if runEnv == "DOCKER" {
		paths = "go/bin/public/*.html"
	} else {
		paths = "public/*.html"
	}
	renderer.NewTemplateRenderer(e, paths) // localhost
	e.GET("/", handleGetAuthor)
	e.GET("/api/author", handleGetAuthorApi)
	e.Logger.Fatal(e.Start(":8080"))
}
