package exporter

import (
	"encoding/base64"
)

func (e *excelExporter) ToBase64() (base64File string, err error) {
	// pre render
	err = e._preRender()
	if err != nil {
		return
	}

	// write to buffer
	buff, err := e.xlsx.WriteToBuffer()
	if err != nil {
		return
	}

	// convert to base64
	base64File = base64.StdEncoding.EncodeToString(buff.Bytes())

	// close the file
	e.xlsx.Close()

	return
}

func (e *excelExporter) Save(filepath string) (err error) {
	// pre render
	err = e._preRender()
	if err != nil {
		return
	}

	// saving a file
	err = e.xlsx.SaveAs(filepath)
	if err != nil {
		return
	}

	// close the file
	e.xlsx.Close()

	return
}

func (e *excelExporter) _preRender() (err error) {
	for _, sheet := range e.sheets {
		e.SetActiveSheet(sheet.SheetName)

		// add header
		err = e._SetHeader(sheet.ConfigFields)
		if err != nil {
			return
		}

		// add style if needed
		if sheet.WithStyle {
			err = e.AddStyle()
			if err != nil {
				return
			}
		}

		// add filter if needed
		if sheet.WithFilter {
			err = e.AddFilter()
			if err != nil {
				return
			}
		}
	}

	return
}
