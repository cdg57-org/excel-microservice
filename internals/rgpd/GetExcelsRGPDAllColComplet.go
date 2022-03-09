package rgpd

import (
	"bytes"
	"excel-microservice/internals/database"
	"excel-microservice/internals/utils"
	"log"
	"reflect"

	"github.com/xuri/excelize/v2"
)

func GetExcelsColComplet() (buf *bytes.Buffer) {
	RGPD_COLLUMS := []string{"RGPD_COL_CODE", "COL_IDENTITE", "COL_EMAIL", "COL_TEL"}
	CDG57s := database.GetRGPDCOMPLET()

	f := excelize.NewFile()

	for i, rgpd_collum := range RGPD_COLLUMS {
		CellID, _ := utils.GetAxis(1, i+1)
		f.SetCellValue(sheet, CellID, rgpd_collum)
	}
	for r, row := range CDG57s {
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
