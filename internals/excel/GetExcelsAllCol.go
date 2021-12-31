package excel

import (
	"bytes"
	"excel-microservice/internals/database"
	"excel-microservice/internals/utils"
	"log"
	"reflect"

	"github.com/xuri/excelize/v2"
)

var (
	err   error
	sheet = "Sheet1"
)

func GetExcelsAllCol() (buf *bytes.Buffer) {
	
	// create slice of the collum name in the excel documents
	RGPD_COLLUMS := []string{"RGPD_COL_CODE", "COL_IDENTITE", "COL_EMAIL", "COL_TEL", "CNRACL", "RG", "AUTRE", "TOTAL", "FACTURE_1ER_ANNEE", "FACTURE_2EME_ANNEE", "FACTURE_3EME_ANNEE", "FACTURE_4EME_ANNEE", "FACTURE_5EME_ANNEE"}
	
	// get data from the sql-server
	CDG57s := database.GetRGPDCOLL()

	// create excel file
	f := excelize.NewFile()


	// insert collum name into the excel 
	for i, rgpd_collum := range RGPD_COLLUMS {
		CellID, _ := utils.GetAxis(1, i+1)
		f.SetCellValue(sheet, CellID, rgpd_collum)
	}

	//insert data into the excel
	for r, row := range CDG57s {
		_ = r
		z := r + 1
		v := reflect.ValueOf(row)


		for i := 0; i < v.NumField(); i++ {
			CellID, _ := utils.GetAxis(z+1, i+1)
			f.SetCellValue(sheet, CellID, v.Field(i).Interface())

		}

	}


	// write the excel to buff for use in the api
	buf, err = f.WriteToBuffer()
	if err != nil {
		log.Println(err)
	}

	//return the buffer
	return buf
}
