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

	sheet.configFields = make([]Field, 0)
	for _, field := range configFields {
		sheet.configFields = append(sheet.configFields, Field{
			Key:                field.Key,
			Label:              field.Label,
			As:                 field.As,
			Default:            field.Default,
			DateFormat:         field.DateFormat,
			DateFormatLocation: field.DateFormatLocation,
			DateParseLayout:    field.DateParseLayout,
			DateParseLocation:  field.DateParseLocation,
			LongestChar:        0,
		})
	}

	e.sheets[e.workingSheetName] = sheet

	return
}

func (e *excelExporter) _SetHeader(configFields []Field) (err error) {
	sheet, ok := e.sheets[e.workingSheetName]
	if !ok {
		return errors.New("sheet not found")
	}

	sheet.configFields = configFields

	// header
	for i, fieldConfig := range sheet.configFields {
		// get cell name
		col, _ := excelize.CoordinatesToCellName(i+1, 1)

		// set value header
		e.xlsx.SetCellValue(e.workingSheetName, col, fieldConfig.Label)

		// set longest char
		if len(fieldConfig.Label) > fieldConfig.LongestChar {
			sheet.configFields[i].LongestChar = len(fieldConfig.Label)
		}
	}

	e.sheets[e.workingSheetName] = sheet

	return
}

func (e *excelExporter) AddRow(mapData map[string]any) (err error) {
	sheet, ok := e.sheets[e.workingSheetName]
	if !ok {
		return errors.New("sheet not found")
	}

	if len(sheet.configFields) == 0 {
		return errors.New("header not found")
	}

	rowValues := make([]any, 0)
	for _, fieldConfig := range sheet.configFields {
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
	for i, fieldConfig := range sheet.configFields {
		col, _ := excelize.CoordinatesToCellName(i+1, (2 + sheet.currentRow))
		val, _ := e.xlsx.GetCellValue(e.workingSheetName, col)

		if len(val) > fieldConfig.LongestChar {
			sheet.configFields[i].LongestChar = len(val)
		}
	}

	sheet.currentRow++

	e.sheets[e.workingSheetName] = sheet

	return
}
