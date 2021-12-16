package main

import (
	"bytes"
	"excel-microservice/internals/excel"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("go-excel-microservice: ")
	log.SetOutput(os.Stderr)

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/", func(c echo.Context) error {
	// 	return c.File("index.html")
	// })
	e.GET("/RGPD_EXPORT", func(c echo.Context) error {
		file := excel.GetExcelsAllCol()
		f := bytes.NewReader(file.Bytes())
		return c.Stream(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", f)
	})
	e.GET("/RGPD_EXPORT_DOSSIER_COMPLET", func(c echo.Context) error {
		file := excel.GetExcelsColComplet()
		f := bytes.NewReader(file.Bytes())
		return c.Stream(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", f)
	})

	e.Logger.Fatal(e.Start(":8013"))

}
