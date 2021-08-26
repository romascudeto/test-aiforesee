package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	Config "aiforesee/config"
	"aiforesee/fuel"
	"aiforesee/fuel/services"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var err error

func main() {
	godotenv.Load()
	Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}
	services.GenerateDataFuel()
	e := echo.New()
	fuel.Routes(e)
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":1323"))
}
