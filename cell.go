package exporter

import (
	"errors"
	"fmt"

	"github.com/faytranevozter/simple-exporter/config"
	"github.com/xuri/excelize/v2"
)

func (e *excelExporter) SetSheetHeader(configFields []config.FieldConfig) (err error) {
	sheet, ok := e.sheets[e.workingSheetName]
	if !ok {
		return errors.New("sheet not found")
	}
	sheet.ConfigFields = configFields
	e.sheets[e.workingSheetName] = sheet

	return
}

func (e *excelExporter) _SetHeader(configFields []config.FieldConfig) (err error) {

	sheet := e.sheets[e.workingSheetName]
	sheet.ConfigFields = configFields

	// header
	for i, fieldConfig := range sheet.ConfigFields {
		// get cell name
		col, _ := excelize.CoordinatesToCellName(i+1, 1)

		// set value header
		e.xlsx.SetCellValue(e.workingSheetName, col, fieldConfig.Label)

		// set longest char
		if len(fieldConfig.Label) > fieldConfig.LongestChar {
			sheet.ConfigFields[i].LongestChar = len(fieldConfig.Label)
		}
	}

	e.sheets[e.workingSheetName] = sheet

	return
}

func (e *excelExporter) AddRow(mapData map[string]any) (err error) {

	sheet := e.sheets[e.workingSheetName]

	rowValues := make([]any, 0)
	for _, fieldConfig := range sheet.ConfigFields {
		var cellValue any

		// get from mapData
		value, ok := mapData[fieldConfig.Key]
		if !ok {
			cellValue = fieldConfig.Default
		} else {
			cellValue = e.Cast(value, fieldConfig)
		}

		// cast cell value
		rowValues = append(rowValues, cellValue)
	}

	// set row value
	e.xlsx.SetSheetRow(e.workingSheetName, fmt.Sprintf("A%d", (2+sheet.currentRow)), &rowValues)

	// set longest char
	for i, fieldConfig := range sheet.ConfigFields {
		col, _ := excelize.CoordinatesToCellName(i+1, (2 + sheet.currentRow))
		val, _ := e.xlsx.GetCellValue(e.workingSheetName, col)

		if len(val) > fieldConfig.LongestChar {
			sheet.ConfigFields[i].LongestChar = len(val)
		}
	}

	sheet.currentRow++

	e.sheets[e.workingSheetName] = sheet

	return
}
