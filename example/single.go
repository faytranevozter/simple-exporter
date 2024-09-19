package main

import (
	exporter "github.com/faytranevozter/simple-exporter"
	"github.com/faytranevozter/simple-exporter/config"
)

func SingleSheetExport() {
	exp := exporter.NewExporter(
		config.WithSheetHeader(
			[]config.FieldConfig{
				{
					Key:   "id",
					Label: "ID",
				},
				{
					Key:   "name",
					Label: "Name",
				},
			},
		),
		config.WithSheetName("List"),
		config.WithSheetFilter(true),
		config.WithSheetStyle(true),
	)

	exp.AddRow(RandomDataset())
	exp.AddRow(RandomDataset())

	exp.Save("single_sheet.xlsx")
}
