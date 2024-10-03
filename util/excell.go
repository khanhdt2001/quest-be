package util

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GenerateFile() (err error) {

	f := excelize.NewFile()
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet2")
	if err != nil {
		return fmt.Errorf("error creating new sheet: %v", err)
	}
	f.SetCellValue("Sheet2", "A1", "Hello world.")
	f.SetActiveSheet(index)
	if err = f.SaveAs("Book1.xlsx"); err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return err
}
