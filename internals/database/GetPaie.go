package database

import (
	"database/sql"
	"excel-microservice/internals/models"
	"fmt"
	"log"
	"os"
)

var (
	PAIE_SQL string = `SELECT  DISTINCT  etudepaye_col_code, COL_IDENTITE, COL_EMAIL, COL_TEL, etudepaye_nbpaye, etudepaye_mission, etudepaye_commentaire
	FROM [data].dbo.etudepaye e , [data].dbo.COLLECTIVITES c 
	WHERE COL_CODE = etudepaye_col_code`
)

func GetPaie() (Paies []models.Etudepaye) {

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, pass, url, database)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Println(err)
	}
	rows, err := db.Query(PAIE_SQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		paye := new(models.Etudepaye)
		if err := rows.Scan(&paye.EtudepayeColCode, &paye.ColIdentite, &paye.ColEmail, &paye.ColTel, &paye.EtudepayeNbpaye, &paye.EtudepayeMission, &paye.EtudepayeCommentaire); err != nil {
			log.Panic(err)
		}
		Paies = append(Paies, *paye)
	}
	if err := rows.Err(); err != nil {
		log.Panic(err)
	}
	return Paies
}
