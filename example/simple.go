package main

import (
	exporter "github.com/faytranevozter/simple-exporter"
	"github.com/faytranevozter/simple-exporter/config"
)

func SimpleExport() {
	exp := exporter.NewExporter()

	// set header
	exp.SetSheetHeader([]config.FieldConfig{
		{
			Key:   "id",
			Label: "ID",
		},
		{
			Key:   "name",
			Label: "Name",
		},
	})

	exp.AddRow(RandomDataset())
	exp.AddRow(RandomDataset())

	exp.Save("simple_sheet.xlsx")
}
