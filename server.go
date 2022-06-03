package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testserver/MySqlDB"

	"github.com/labstack/echo/v4"
)

type MyConfig struct {
	ServerPort string      `json:"server_port"`
	MYSql      MySQLConfig `json:"mysql"`
}
type MySQLConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func testdb(c echo.Context) error {
	stat, comment := MySqlDB.TestDB()

	if stat {
		return c.String(http.StatusOK, "MySql DB is Connected. Port number is "+comment)
	} else {
		return c.String(http.StatusOK, "MySql DB is not Connected")
	}
}

func insertdb(c echo.Context) error {
	stat, num := MySqlDB.InsertDB()

	if stat {
		return c.String(http.StatusOK, "Insert OK. Row count is "+num)
	} else {
		return c.String(http.StatusOK, "Error in insert data")
	}
}

func main() {

	var PNum string

	// #region Load confgi from .env file
	/*
		var appConfig map[string]string
		appConfig, err := godotenv.Read()

		if err != nil {
			log.Fatal("Error reading .env file")
		}
		PNum = ":" + appConfig["SERVER_PORT"]
	*/
	// #endregion

	// Load config from CONFIG.json file
	file, _ := ioutil.ReadFile("CONFIG.json")
	data := MyConfig{}
	_ = json.Unmarshal([]byte(file), &data)

	PNum = ":" + data.ServerPort
	fmt.Println("Server is running on port", PNum)

	// Create and start a server with created config
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server_OK")
	})

	e.GET("/testdb", testdb)
	e.GET("/insert", insertdb)
	e.Logger.Fatal(e.Start(PNum))

}
