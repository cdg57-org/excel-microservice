package main

import (
	"bytes"
	"excel-microservice/internals/excel"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// _ = file
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	// e.Logger.Fatal(e.Start("127.0.0.1:8013"))
	if err := e.Start(":1337"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
