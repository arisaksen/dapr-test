package main

import (
	"github.com/arisaksen/dapr-test/state_management/api1/author"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	httpClient         http.Client
	daprHost           string
	daprHttpPort       string
	runtimeEnvironment string
)

const (
	stateStoreComponentName = "statestore"
	defaultDaprPort         = "3500"
)

func getAuthor(client *http.Client, name string) (author.Author, error) {
	getResponse, err := client.Get(daprHost + ":" + daprHttpPort + "/v1.0/state/" + stateStoreComponentName + "/" + name)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(getResponse.Body)
	if err != nil {
		panic(err)
	}
	valueAsString := string(body)
	valueAsStringTrimmed := strings.ReplaceAll(valueAsString, "\"", "")
	valueAsInt, err := strconv.Atoi(valueAsStringTrimmed)
	if err != nil {
		panic(err)
	}
	a := author.Author{
		Name:        name,
		YearOfBirth: valueAsInt,
	}

	return a, nil
}

func handleGetAuthor(c echo.Context) error {
	name := c.Param("name")
	authors, err := getAuthor(&httpClient, name)
	if err != nil {
		panic(err)
	}
	response := c.JSON(http.StatusOK, authors)
	return response
}

func init() {
	runtimeEnvironment = os.Getenv("ENVIRONMENT")
	if runtimeEnvironment == "" {
		runtimeEnvironment = "LOCALHOST"
	}

	httpClient = http.Client{
		Timeout: 15 * time.Second,
	}

	// DAPR path
	daprHost = os.Getenv("DAPR_HOST")
	if daprHost == "" {
		daprHost = "http://localhost"
	}
	daprHttpPort = os.Getenv("DAPR_HTTP_PORT")
	if daprHttpPort == "" {
		daprHttpPort = defaultDaprPort
	}

}

func main() {
	e := echo.New()
	e.Use(responseTimeLogger)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/api/author/:name", handleGetAuthor)
	e.Logger.Fatal(e.Start(":8081"))
}

// example data
//authors := []author.Author{
//	{Name: "J.R.R Tolkien", YearOfBirth: 1892},
//	{Name: "Dan Brown", YearOfBirth: 1964},
//	{Name: "Sun Tzu", YearOfBirth: -544},
//}

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
