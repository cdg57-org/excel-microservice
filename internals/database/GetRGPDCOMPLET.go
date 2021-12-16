package database

import (
	"database/sql"
	"excel-microservice/internals/models"
	"fmt"
	"log"
	"os"
)

var (
	RGPD_COL_SQL = "SELECT RGPD_COL_CODE, COL_IDENTITE, COL_EMAIL, COL_TEL FROM DATA.DBO.RGPD, DATA.DBO.COLLECTIVITES WHERE COL_CODE = rgpd_col_code AND RGPD_CONVENTION = 1 AND RGPD_DELIBERATION = 1 AND RGPD_LETTRE_DE_MISSION = 1"
)

func GetRGPDCOMPLET() (CDG57s []models.RGPD_COLL_COMPLET) {

	// github.com/denisenkom/go-mssqldb
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("sqlserver://%s:%s@srv-application?database=%s", user, pass, database)
	// dsn := ""
	// dquery := url.Values{}
	// query.Add("app name", "MyAppName")

	// }
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Println(err)
	}
	rows, err := db.Query(RGPD_COL_SQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// var
	// log.Println(rows.Columns())
	for rows.Next() {
		cdg57 := new(models.RGPD_COLL_COMPLET)
		if err := rows.Scan(&cdg57.RgpdColCode, &cdg57.ColIdentite, &cdg57.ColEmail, &cdg57.ColTel); err != nil {
			log.Panic(err)
		}
		CDG57s = append(CDG57s, *cdg57)
	}
	if err := rows.Err(); err != nil {
		log.Panic(err)
	}
	return CDG57s

}
