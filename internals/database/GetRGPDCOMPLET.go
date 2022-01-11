package database

import (
	"database/sql"
	"excel-microservice/internals/models"
	"fmt"
	"log"
	"os"
)

var (

	// sql variable to query the collectivité who responded with every documents 
	RGPD_COL_SQL = `SELECT 
	RGPD_COL_CODE, 
	COL_IDENTITE, 
	COL_EMAIL, 
	COL_TEL 
  FROM 
	DATA.DBO.RGPD, 
	DATA.DBO.COLLECTIVITES 
  WHERE 
	COL_CODE = rgpd_col_code 
	AND RGPD_CONVENTION = 1 
	AND RGPD_DELIBERATION = 1 
	AND RGPD_LETTRE_DE_MISSION = 1
  `
)

func GetRGPDCOMPLET() (CDG57s []models.RGPD_COLL_COMPLET) {

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, pass, url, database)



	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Println(err)
	}
	rows, err := db.Query(RGPD_COL_SQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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
