package main

import (
	"time"

	"github.com/go-faker/faker/v4"
)

func main() {
	SimpleExport()
	SingleSheetExport()
	MultiSheetExport()
	TimeExport()
}

func RandomDataset() map[string]any {
	timeString := faker.Timestamp()
	t, _ := time.Parse("2006-01-02 15:04:05", timeString)
	return map[string]any{
		"id":                faker.UUIDHyphenated(),
		"name":              faker.Name(),
		"address":           faker.GetRealAddress(),
		"created_at_string": timeString,
		"created_at_time":   t,
	}
}
