package config

func DefaultMultipleConfig() OptSheet {
	return OptSheet{
		SheetOpt: []Opts{
			DefaultConfig(),
		},
		withStyle:  false,
		withFilter: false,
	}
}

func WithSheetHeaders(headers ...[]FieldConfig) func(*OptSheet) {
	return func(o *OptSheet) {
		if len(o.SheetOpt) == 0 {
			o.SheetOpt = make([]Opts, len(headers))
		}

		for i, fieldConfigs := range headers {
			if len(o.SheetOpt) < (i + 1) {
				o.SheetOpt = append(o.SheetOpt, DefaultConfig())
			}

			o.SheetOpt[i].ConfigFields = fieldConfigs
		}
	}
}

func WithSheetNames(sheetNames ...string) func(*OptSheet) {
	return func(o *OptSheet) {
		if len(o.SheetOpt) == 0 {
			o.SheetOpt = make([]Opts, len(sheetNames))
		}

		for i, sheetName := range sheetNames {
			if len(o.SheetOpt) < (i + 1) {
				o.SheetOpt = append(o.SheetOpt, DefaultConfig())
			}

			o.SheetOpt[i].SheetName = sheetName
		}
	}
}

func WithSheetStyles(enables ...bool) func(o *OptSheet) {
	return func(o *OptSheet) {
		if len(o.SheetOpt) == 0 {
			o.SheetOpt = make([]Opts, len(enables))
		}

		for i, enabled := range enables {
			if len(o.SheetOpt) < (i + 1) {
				o.SheetOpt = append(o.SheetOpt, DefaultConfig())
			}

			o.SheetOpt[i].WithStyle = enabled
		}
	}
}

func WithSheetFilters(enables ...bool) func(o *OptSheet) {
	return func(o *OptSheet) {
		if len(o.SheetOpt) == 0 {
			o.SheetOpt = make([]Opts, len(enables))
		}

		for i, enabled := range enables {
			if len(o.SheetOpt) < (i + 1) {
				o.SheetOpt = append(o.SheetOpt, DefaultConfig())
			}

			o.SheetOpt[i].WithFilter = enabled
		}
	}
}
