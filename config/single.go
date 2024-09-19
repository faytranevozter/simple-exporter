package config

import "fmt"

func DefaultConfig(i int) Opts {
	return Opts{
		ConfigFields: []FieldConfig{},
		WithStyle:    false,
		SheetName:    fmt.Sprintf("Sheet%d", i+1),
	}
}

func WithSheetHeader(headers []FieldConfig) func(*Opts) {
	return func(o *Opts) {
		o.ConfigFields = headers
	}
}

func WithSheetName(sheetName string) func(*Opts) {
	return func(o *Opts) {
		o.SheetName = sheetName
	}
}

func WithSheetStyle(enable bool) func(o *Opts) {
	return func(o *Opts) {
		o.WithStyle = enable
	}
}

func WithSheetFilter(enable bool) func(o *Opts) {
	return func(o *Opts) {
		o.WithFilter = enable
	}
}
