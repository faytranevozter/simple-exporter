package main

import (
	exporter "github.com/faytranevozter/simple-exporter"
	"github.com/faytranevozter/simple-exporter/config"
)

func MultiSheetExport() {
	exp := exporter.NewExporterMultiSheet(
		config.WithSheetHeaders(
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
			[]config.FieldConfig{
				{
					Key:   "id",
					Label: "ID 2",
				},
				{
					Key:   "name",
					Label: "Name 2",
				},
				{
					Key:        "created_at",
					Label:      "Created At",
					As:         "date",
					Default:    "-",
					DateFormat: "DATETIME",
				},
			},
		),
		config.WithSheetNames("List 1", "List 2"),
	)

	exp.SetActiveSheetIndex(0)
	exp.AddRow(RandomDataset())

	exp.SetActiveSheetIndex(1)
	exp.SetSheetStyle(true)
	exp.AddRow(RandomDataset())
	exp.AddRow(RandomDataset())
	exp.AddRow(RandomDataset())

	exp.Save("multi_sheet.xlsx")
}
