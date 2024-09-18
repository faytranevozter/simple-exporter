package exporter

import (
	"errors"

	"github.com/faytranevozter/simple-exporter/config"
)

func (e *excelExporter) AddSheet(sheetName string, setActive bool) (err error) {
	idx, err := e.xlsx.NewSheet(sheetName)
	if err != nil {
		return
	}

	e.sheets[sheetName] = Sheet{
		index: idx,
		Opts: config.Opts{
			ConfigFields: []config.FieldConfig{},
			WithStyle:    false,
			WithFilter:   false,
			SheetName:    sheetName,
		},
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

	sheet.SheetName = newSheetName
	e.sheets[newSheetName] = sheet
	e.workingSheetName = newSheetName

	// remove old map
	delete(e.sheets, oldName)

	return
}
