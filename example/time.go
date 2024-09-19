package main

import (
	"time"

	exporter "github.com/faytranevozter/simple-exporter"
	"github.com/faytranevozter/simple-exporter/config"
)

func TimeExport() {
	exp := exporter.NewExporter(
		config.WithSheetStyle(true),
	)

	asiaJakarta, _ := time.LoadLocation("Asia/Jakarta")

	// set header
	exp.SetSheetHeader([]config.FieldConfig{
		{
			Key:   "created_at_string",
			Label: "Created At Original",
		},
		{
			Key:                "created_at_string",
			Label:              "Created At Parse",
			As:                 "date",
			Default:            "-",
			DateParseLayout:    "2006-01-02 15:04:05",
			DateParseLocation:  time.UTC,
			DateFormat:         "DATETIME",
			DateFormatLocation: asiaJakarta,
		},
		{
			Key:                "created_at_time",
			Label:              "Created At",
			As:                 "date",
			Default:            "-",
			DateFormat:         "DATETIME",
			DateFormatLocation: asiaJakarta,
		},
	})

	for i := 0; i < 5; i++ {
		exp.AddRow(RandomDataset())
	}

	exp.Save("time_sheet.xlsx")
}
