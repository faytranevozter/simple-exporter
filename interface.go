package exporter

import (
	"time"

	"github.com/faytranevozter/simple-exporter/config"
	"github.com/xuri/excelize/v2"
)

type Exporter interface {
	SetSheetHeader(headers []config.FieldConfig) (err error)
	SetSheetStyle(enabled bool) (err error)
	SetSheetFilter(enabled bool) (err error)
	SetActiveSheet(sheetName string) (err error)
	SetActiveSheetIndex(index int) (err error)
	AddRow(mapData map[string]any) (err error)
	AddSheet(sheetName string, setActive bool) (err error)
	RenameSheet(newSheetName string) (err error)
	ToBase64() (base64File string, err error)
	Save(filepath string) error
}

type Sheet struct {
	index        int
	currentRow   int
	configFields []Field
	withStyle    bool
	withFilter   bool
	sheetName    string
}

type Field struct {
	Key                string
	Label              string
	As                 string
	Default            string
	DateFormat         string
	DateFormatLocation *time.Location
	DateParseLayout    string
	DateParseLocation  *time.Location
	LongestChar        int
}

type excelExporter struct {
	xlsx              *excelize.File
	workingSheetName  string
	workingSheetIndex int
	sheets            map[string]Sheet
}
