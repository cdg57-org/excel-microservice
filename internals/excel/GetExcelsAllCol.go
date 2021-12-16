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
	addr  string
	sheet = "Sheet1"
)

func GetExcelsAllCol() (buf *bytes.Buffer) {
	RGPD_COLLUMS := []string{"RGPD_COL_CODE", "COL_IDENTITE", "COL_EMAIL", "COL_TEL", "CNRACL", "RG", "AUTRE", "TOTAL", "FACTURE_1ER_ANNEE", "FACTURE_2EME_ANNEE", "FACTURE_3EME_ANNEE", "FACTURE_4EME_ANNEE", "FACTURE_5EME_ANNEE"}
	CDG57s := database.GetRGPDCOLL()
	_ = CDG57s
	// for _, coll := range CDG57s {
	// 	log.Println(coll)
	// }

	f := excelize.NewFile()

	for i, rgpd_collum := range RGPD_COLLUMS {
		CellID, _ := utils.GetAxis(1, i)
		// log.Println(, rgpd_collum)
		f.SetCellValue(sheet, CellID, rgpd_collum)
	}
	for r, row := range CDG57s {
		_ = r
		z := r + 1
		// log.Println(z, row)
		v := reflect.ValueOf(row)
		// log.Println(v.NumField())

		// values := make([]interface{}, v.NumField())

		for i := 0; i < v.NumField(); i++ {
			// log.Println(r, i)
			CellID, _ := utils.GetAxis(z+1, i)
			// log.Println(CellID)
			f.SetCellValue(sheet, CellID, v.Field(i).Interface())

			// v.Field(i).Interface()
		}

		// if err = f.SetSheetRow(sheet, addr, &row); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}

	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	log.Println(err)
	// }
	buf, err = f.WriteToBuffer()
	if err != nil {
		log.Println(err)
	}
	return buf
}
