package config

import "time"

type FieldConfig struct {
	Key                string
	Label              string
	As                 string
	Default            string
	DateFormat         string
	DateFormatLocation *time.Location
	DateParseLayout    string
	DateParseLocation  *time.Location
	LongestChar        int
}
