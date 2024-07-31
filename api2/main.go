package main

import (
	"encoding/json"
	"github.com/arisaksen/api1/author"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"time"
)

type Book struct {
	Name   string        `json:"name"`
	Author author.Author `json:"author"`
	Year   int           `json:"year"`
}

func GetBooks(authors []author.Author) ([]Book, error) {
	books := []Book{
		Book{Name: "The Hobbit", Author: authors[0], Year: 1937},
		Book{Name: "The Da Vinci Code", Author: authors[1], Year: 2003},
		Book{Name: "The Art of War", Author: authors[2], Year: -475},
	}

	return books, nil
}

func GetAuthors() ([]author.Author, error) {
	response, err := http.Get("http://localhost:8080/api/author")
	if err != nil {
		panic(err.Error())
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	var authors []author.Author
	err = json.Unmarshal(body, &authors)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func handleGetBook(c echo.Context) error {
	authors, err := GetAuthors()
	if err != nil {
		panic(err.Error())
	}
	books, err := GetBooks(authors)
	if err != nil {
		panic(err.Error())
	}

	response := c.JSON(http.StatusOK, books)
	return response
}

func main() {
	e := echo.New()
	e.Use(responseTimeLogger)
	e.GET("/api/book", handleGetBook)
	e.Logger.Fatal(e.Start(":8081"))
}

func responseTimeLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		end := time.Now()
		responseTime := end.Sub(start)

		log.Printf("method=%s, uri=%s, status=%d, response_time=%s",
			c.Request().Method,
			c.Request().RequestURI,
			c.Response().Status,
			responseTime)

		return err
	}
}
