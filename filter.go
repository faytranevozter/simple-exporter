package exporter

import (
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (e *excelExporter) SetSheetFilter(enabled bool) (err error) {
	sheet, ok := e.sheets[e.workingSheetName]
	if !ok {
		return errors.New("sheet not found")
	}

	sheet.withFilter = enabled
	e.sheets[e.workingSheetName] = sheet

	return
}

func (e *excelExporter) AddFilter() (err error) {
	sheet := e.sheets[e.workingSheetName]
	if len(sheet.configFields) == 0 {
		return
	}

	lastCol, _ := excelize.ColumnNumberToName(len(sheet.configFields))
	rangeRef := fmt.Sprintf("A1:%s1", lastCol)

	err = e.xlsx.AutoFilter(e.workingSheetName, rangeRef, []excelize.AutoFilterOptions{})
	if err != nil {
		return
	}

	return
}
