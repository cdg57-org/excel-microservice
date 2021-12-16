package database

import (
	"database/sql"
	"excel-microservice/internals/models"
	"fmt"
	"log"
	"os"
)

var (
	RGPD_COL string = `SELECT DISTINCT RGPD_COL_CODE, COL_IDENTITE, COL_EMAIL, COL_TEL, sum(COT_ASS1) AS 'CNRACL', sum(COT_ASS2) AS 'RG', sum(COT_ASS3) AS 'AUTRE', sum(COT_ASS1+COT_ASS2+COT_ASS3) AS 'TOTAL', 
	case
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) < 100000 THEN 760 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 100000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 300000  THEN 950 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 300000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 500000  THEN 1050 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 500000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 1000000  THEN 1250 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) >  1000000  THEN 1450 
	 END as FACTURE_1ER_ANNEE, 
	 case
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) < 100000 THEN 200 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 100000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 300000  THEN 250 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 300000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 500000  THEN 300 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 500000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 1000000  THEN 350 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) >  1000000  THEN 400 
	 END as FACTURE_2EME_ANNEE,
	 case
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) < 100000 THEN 200 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 100000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 300000  THEN 250 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 300000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 500000  THEN 300 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 500000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 1000000  THEN 350 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) >  1000000  THEN 400 
	 END as FACTURE_3EME_ANNEE, 
	  case
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) < 100000 THEN 200 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 100000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 300000  THEN 250 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 300000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 500000  THEN 300 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 500000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 1000000  THEN 350 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) >  1000000  THEN 400 
	 END as FACTURE_4EME_ANNEE,
	 case
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) < 100000 THEN 200 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 100000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 300000  THEN 250 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 300000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 500000  THEN 300 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) > 500000 AND sum(COT_ASS1+COT_ASS2+COT_ASS3) < 1000000  THEN 350 
	 WHEN  sum(COT_ASS1+COT_ASS2+COT_ASS3) >  1000000  THEN 400 
	 END as FACTURE_5EME_ANNEE
	 FROM DATA.DBO.RGPD, DATA.DBO.COLLECTIVITES, DATA.DBO.COTISATIONS
	 WHERE COL_CODE = rgpd_col_code and COL_ID = COT_COLID 
	 and COT_ANNEE = 2020 
	 GROUP BY RGPD_COL_CODE, COL_IDENTITE, COL_EMAIL, COL_TEL`
)

func GetRGPDCOLL() (CDG57s []models.RGPD_COLL) {
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
	rows, err := db.Query(RGPD_COL) //db.Query("SELECT RGPD_COL_CODE, COL_IDENTITE, COL_EMAIL, COL_TEL FROM DATA.DBO.RGPD, DATA.DBO.COLLECTIVITES WHERE COL_CODE = rgpd_col_code AND RGPD_CONVENTION = 1 AND RGPD_DELIBERATION = 1 AND RGPD_LETTRE_DE_MISSION = 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// var
	// log.Println(rows.Columns())
	for rows.Next() {
		cdg57 := new(models.RGPD_COLL)
		if err := rows.Scan(&cdg57.RgpdColCode, &cdg57.ColIdentite, &cdg57.ColEmail, &cdg57.ColTel, &cdg57.Cnracl, &cdg57.Rg, &cdg57.Autre, &cdg57.Total, &cdg57.Facture1ErAnnee, &cdg57.Facture2EmeAnnee, &cdg57.Facture3EmeAnnee, &cdg57.Facture4EmeAnnee, &cdg57.Facture5EmeAnnee); err != nil {
			log.Panic(err)
		}
		CDG57s = append(CDG57s, *cdg57)
	}
	if err := rows.Err(); err != nil {
		log.Panic(err)
	}
	return CDG57s
}
