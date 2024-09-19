package exporter

import (
	"errors"
)

func (e *excelExporter) AddSheet(sheetName string, setActive bool) (err error) {
	idx, err := e.xlsx.NewSheet(sheetName)
	if err != nil {
		return
	}

	e.sheets[sheetName] = Sheet{
		index:        idx,
		configFields: []Field{},
		withStyle:    false,
		withFilter:   false,
		sheetName:    sheetName,
	}

	if setActive {
		e.SetActiveSheet(sheetName)
	}

	return
}

func (e *excelExporter) SetActiveSheet(sheetName string) (err error) {
	sheet, ok := e.sheets[sheetName]
	if !ok {
		return errors.New("sheet not found")
	}

	e.xlsx.SetActiveSheet(sheet.index)
	e.workingSheetName = sheetName
	e.workingSheetIndex = sheet.index

	return
}

func (e *excelExporter) SetActiveSheetIndex(index int) (err error) {
	sheet := Sheet{}
	for _, sh := range e.sheets {
		if sh.index == index {
			sheet = sh
		}
	}

	if sheet.sheetName == "" {
		return errors.New("sheet not found")
	}

	e.xlsx.SetActiveSheet(sheet.index)
	e.workingSheetName = sheet.sheetName
	e.workingSheetIndex = sheet.index

	return
}

func (e *excelExporter) RenameSheet(newSheetName string) (err error) {
	oldName := e.workingSheetName

	if oldName == newSheetName {
		return errors.New("new sheet name is same as old sheet name")
	}

	sheet, ok := e.sheets[oldName]
	if !ok {
		return errors.New("sheet not found")
	}

	err = e.xlsx.SetSheetName(oldName, newSheetName)
	if err != nil {
		return err
	}

	sheet.sheetName = newSheetName
	e.sheets[newSheetName] = sheet
	e.workingSheetName = newSheetName

	// remove old map
	delete(e.sheets, oldName)

	return
}
