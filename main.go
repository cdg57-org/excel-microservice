package main

import (
	"bytes"
	"excel-microservice/internals/paie"
	"excel-microservice/internals/rgpd"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/paie", fmt.Sprintf("%s/public", os.Getenv("PAIE_FOLDER")))

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

	e.POST("/NewPaie_Export", func(c echo.Context) error {

		file := paie.GetNewExcelPaie()
		f := bytes.NewReader(file.Bytes())
		return c.Stream(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", f)
	})

	e.POST("/PaieUpload", upload)

	if len(os.Getenv("PAIE_FOLDER")) == 0 {
		e.GET("/PaieDL", func(c echo.Context) error {
			return c.File("downloads/paie.xlsx")
		})
	} else {
		e.GET("/PaieDL", func(c echo.Context) error {
			return c.File(fmt.Sprintf("%s/downloads/paie.xlsx", os.Getenv("PAIE_FOLDER")))
		})
	}

	if err := e.Start(os.Getenv("ADDR")); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func upload(c echo.Context) error {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username

	log.Printf("Username: %s\n", username)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if len(os.Getenv("PAIE_FOLDER")) == 0 {
		dst, err := os.Create(fmt.Sprintf("downloads/%s", "paie.xlsx"))
		// Destination

		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	} else {
		dst, err := os.Create(fmt.Sprintf("%s/downloads/paie.xlsx", os.Getenv("PAIE_FOLDER")))
		// Destination

		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
}
