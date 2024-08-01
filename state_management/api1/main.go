package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/arisaksen/dapr-test/state_management/api1/author"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
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

func postAuthor(client *http.Client, a author.Author) {
	var state []byte
	state, _ = json.Marshal([]map[string]string{
		{
			"key":   a.Name,
			"value": strconv.Itoa(a.YearOfBirth),
		},
	})
	res, err := client.Post(daprHost+":"+daprHttpPort+"/v1.0/state/"+stateStoreComponentName, "application/json", bytes.NewReader(state))
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Println("Saved Author:", a)
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

func handlePostAuthor(c echo.Context) error {
	var body author.Author
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		log.Error("empty json body")
	}
	postAuthor(&httpClient, body)

	response := c.JSON(http.StatusCreated, body)
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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/api/author/:name", handleGetAuthor)
	e.POST("/api/author", handlePostAuthor)
	e.Logger.Fatal(e.Start(":8080"))
}

// example data
//authors := []author.Author{
//	{Name: "J.R.R Tolkien", YearOfBirth: 1892},
//	{Name: "Dan Brown", YearOfBirth: 1964},
//	{Name: "Sun Tzu", YearOfBirth: -544},
//}
