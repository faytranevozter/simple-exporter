package exporter

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

func (e *excelExporter) SetSheetStyle(enabled bool) (err error) {
	sheet, ok := e.sheets[e.workingSheetName]
	if !ok {
		return errors.New("sheet not found")
	}

	sheet.WithStyle = enabled
	e.sheets[e.workingSheetName] = sheet

	return
}

func (e *excelExporter) AddStyle() (err error) {
	sheet := e.sheets[e.workingSheetName]

	style, errStyle := e.xlsx.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 15,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E0EBF5"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if errStyle != nil {
		err = errStyle
		return
	}

	lastHeaderCol, _ := excelize.CoordinatesToCellName(len(sheet.ConfigFields), 1)
	e.xlsx.SetCellStyle(e.workingSheetName, "A1", lastHeaderCol, style)

	// set column width
	for i, fieldConfig := range sheet.ConfigFields {
		// get cell name
		colName, _ := excelize.ColumnNumberToName(i + 1)
		e.xlsx.SetColWidth(e.workingSheetName, colName, colName, float64(fieldConfig.LongestChar)*1.2)
	}

	return
}
