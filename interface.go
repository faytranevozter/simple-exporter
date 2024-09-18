package exporter

import (
	"github.com/faytranevozter/simple-exporter/config"
	"github.com/xuri/excelize/v2"
)

type Exporter interface {
	SetSheetHeader(headers []config.FieldConfig) (err error)
	SetSheetStyle(enabled bool) (err error)
	SetSheetFilter(enabled bool) (err error)
	SetActiveSheet(sheetName string) (err error)
	AddRow(mapData map[string]any) (err error)
	AddSheet(sheetName string, setActive bool) (err error)
	RenameSheet(newSheetName string) (err error)
	ToBase64() (base64File string, err error)
	Save(filepath string) error
}

type Sheet struct {
	index      int
	currentRow int
	config.Opts
}

type excelExporter struct {
	xlsx              *excelize.File
	workingSheetName  string
	workingSheetIndex int
	sheets            map[string]Sheet
}
