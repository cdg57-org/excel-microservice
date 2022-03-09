package main

import (
	"bytes"
	"excel-microservice/internals/paie"
	"excel-microservice/internals/rgpd"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/labstack/echo/v4"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("go-excel-microservice: ")
	log.SetOutput(os.Stderr)

}

func main() {

	switch ostype := runtime.GOOS; ostype {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		// fmt.Println("Linux.")
		// err := godotenv.Load("/etc/excel-microservice/config.env", "./config.env")
		// if err != nil {
		// 	log.Fatal("Error loading .env file")
		// }

		env := os.Getenv("ENV")
		if env == "dev" {
			err := godotenv.Load("config.env")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := godotenv.Load("/etc/excel-microservice/config.env")
			if err != nil {
				log.Fatal(err)
			}
		}
	case "windows":
		err := godotenv.Load("config.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	default:
		log.Printf("%s.\n", ostype)
	}

	log.Println(runtime.GOOS)

	// _ = file
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// RGPD
	e.POST("/RGPD_EXPORT", func(c echo.Context) error {
		file := rgpd.GetExcelsAllCol()
		f := bytes.NewReader(file.Bytes())
		return c.Stream(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", f)
	})
	e.POST("/RGPD_EXPORT_DOSSIER_COMPLET", func(c echo.Context) error {
		file := rgpd.GetExcelsColComplet()
		f := bytes.NewReader(file.Bytes())
		return c.Stream(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", f)
	})

	// PAIE

	e.POST("/PAIE_REPORT", func(c echo.Context) error {
		file := paie.GetExcelPaie()
		f := bytes.NewReader(file.Bytes())
		return c.Stream(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", f)
	})

	// e.Logger.Fatal(e.Start("127.0.0.1:8013"))
	if err := e.Start(os.Getenv("ADDR")); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
