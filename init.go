package exporter

import (
	"github.com/faytranevozter/simple-exporter/config"
	"github.com/xuri/excelize/v2"
)

func NewExporter(opts ...config.OptFunc) Exporter {

	option := config.DefaultConfig()

	for _, fn := range opts {
		fn(&option)
	}

	xlsx := excelize.NewFile()

	// set default sheet name
	xlsx.SetSheetName(xlsx.GetSheetName(0), option.SheetName)

	activeSheetIndex := xlsx.GetActiveSheetIndex()

	return &excelExporter{
		xlsx:              xlsx,
		workingSheetName:  option.SheetName,
		workingSheetIndex: activeSheetIndex,
		sheets: map[string]Sheet{
			option.SheetName: {
				index: activeSheetIndex,
				Opts:  option,
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

		exp.sheets[opt.SheetName] = Sheet{
			index: i,
			Opts:  opt,
		}
	}

	return exp
}
