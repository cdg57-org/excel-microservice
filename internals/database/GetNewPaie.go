package database

import (
	"database/sql"
	"excel-microservice/internals/models"
	"fmt"
	"log"
	"os"
)

var (
	NewPAIE_SQL string = `
SELECT
	DISTINCT etudepaye_col_code,
	COL_IDENTITE,
	etudepaye_nbpaye,
	etudepaye_notif_engagement,
	etudepaye_lettre_engagement,
	etudepaye_deliberation,
	etudepaye_convention
FROM
	DATA.DBO.etudepaye2,
	DATA.DBO.COLLECTIVITES
WHERE
	etudepaye_col_code = COL_CODE
`
)

func GetNewPaie() (NewPaie []models.NewEtudepaye) {
	// log.Println(NewPAIE_SQL)

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, pass, url, database)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Println(err)
	}
	rows, err := db.Query(NewPAIE_SQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		paye := new(models.NewEtudepaye)
		if err := rows.Scan(&paye.EtudepayeId, &paye.EtudepayeColCode, &paye.EtudepayeNbpaye, &paye.EtudepayeNotifEngagement, &paye.EtudepayeLettreEngagement, &paye.EtudepayeDeliberation, &paye.EtudepayeConvention); err != nil {
			log.Panic(err)
		}
		NewPaie = append(NewPaie, *paye)
	}
	if err := rows.Err(); err != nil {
		log.Panic(err)
	}

	return NewPaie
}
