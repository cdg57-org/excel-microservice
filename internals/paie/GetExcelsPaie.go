package paie

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

func GetExcelPaie() (buf *bytes.Buffer) {

	PAIES_COLLUMS := []string{"etudepaye_col_code", "COL_IDENTITE", "COL_EMAIL", "COL_TEL", "etudepaye_nbpaye", "etudepaye_mission", "etudepaye_commentaire"}
	payes := database.GetPaie()

	f := excelize.NewFile()

	for i, rgpd_collum := range PAIES_COLLUMS {
		CellID, _ := utils.GetAxis(1, i+1)
		f.SetCellValue(sheet, CellID, rgpd_collum)
	}
	for r, row := range payes {
		_ = r
		z := r + 1
		v := reflect.ValueOf(row)

		for i := 0; i < v.NumField(); i++ {
			CellID, _ := utils.GetAxis(z+1, i+1)
			f.SetCellValue(sheet, CellID, v.Field(i).Interface())

		}

	}

	buf, err = f.WriteToBuffer()
	if err != nil {
		log.Println(err)
	}
	return buf

}
