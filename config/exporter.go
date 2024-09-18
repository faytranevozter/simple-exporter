package config

type Opts struct {
	ConfigFields []FieldConfig
	WithStyle    bool
	WithFilter   bool
	SheetName    string
}

type OptSheet struct {
	SheetOpt   []Opts
	withStyle  bool
	withFilter bool
}

type OptFunc func(e *Opts)

type SheetOptFunc func(*OptSheet)
