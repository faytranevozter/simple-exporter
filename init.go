package exporter

import (
	"github.com/faytranevozter/simple-exporter/config"
	"github.com/xuri/excelize/v2"
)

func NewExporter(opts ...config.OptFunc) Exporter {

	option := config.DefaultConfig(0)

	for _, fn := range opts {
		fn(&option)
	}

	xlsx := excelize.NewFile()

	// set default sheet name
	xlsx.SetSheetName(xlsx.GetSheetName(0), option.SheetName)

	activeSheetIndex := xlsx.GetActiveSheetIndex()

	fields := make([]Field, 0)
	for _, field := range option.ConfigFields {
		fields = append(fields, Field{
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

	return &excelExporter{
		xlsx:              xlsx,
		workingSheetName:  option.SheetName,
		workingSheetIndex: activeSheetIndex,
		sheets: map[string]Sheet{
			option.SheetName: {
				index:        0,
				configFields: fields,
				withStyle:    option.WithStyle,
				withFilter:   option.WithFilter,
				sheetName:    option.SheetName,
			},
		},
	}
}

func NewExporterMultiSheet(sheetOpts ...config.SheetOptFunc) Exporter {

	option := config.DefaultMultipleConfig()

	for _, fn := range sheetOpts {
		fn(&option)
	}

	xlsx := excelize.NewFile()

	activeSheetIndex := xlsx.GetActiveSheetIndex()

	exp := &excelExporter{
		xlsx:              xlsx,
		workingSheetName:  option.SheetOpt[0].SheetName,
		workingSheetIndex: activeSheetIndex,
		sheets:            make(map[string]Sheet),
	}

	for i, opt := range option.SheetOpt {
		if i == 0 {
			exp.xlsx.SetSheetName(exp.xlsx.GetSheetName(i), opt.SheetName)
		} else {
			exp.xlsx.NewSheet(opt.SheetName)
		}

		fields := make([]Field, 0)
		for _, field := range opt.ConfigFields {
			fields = append(fields, Field{
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

		exp.sheets[opt.SheetName] = Sheet{
			index:        i,
			configFields: fields,
			withStyle:    opt.WithStyle,
			withFilter:   opt.WithFilter,
			sheetName:    opt.SheetName,
		}
	}

	return exp
}
