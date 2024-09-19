# Simple Exporter

A simple exporter for monitoring and metrics collection.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Simple Exporter is a lightweight tool designed to collect and export metrics for monitoring purposes. It is built with Go and can be easily integrated into your existing infrastructure.

## Installation
```bash
go get github.com/faytranevozter/simple-exporter
```

# Usage

### Setup
```go
import (
  exporter "github.com/faytranevozter/simple-exporter"
)

func main() {
  // ...

  // minimal setup
  exp := exporter.NewExporter()

  // or with configuration
  exp := exporter.NewExporter(
    config.WithSheetHeader(
      []config.FieldConfig{
        {
          Key:   "number",
          Label: "No",
        },
        {
          Key:        "created",
          Label:      "Created At",
          As:         "date",
          Default:    "-",
          DateFormat: "DATETIME",
        },
      },
    ),
    config.WithSheetName("List User"),
    config.WithSheetStyle(true),
    config.WithSheetFilter(true),
  )

  // ...
}
```

### Adding Data
```go
for _, row := range someListData {
  exp.AddRow(map[string]any{
    "id":         row.ID,
    "name":       row.Name,
    "age":        row.Age,
    "created_at": row.CreatedAt,
  })
}
```

### Exporting
```go
err := exp.Save("sample.xlsx")
if err != nil {
  log.Fatal(err)
}
```

- [See example](./example/).

## Configuration

FieldConfig for Header:
```go
type FieldConfig struct {
  Key                string
  Label              string
  As                 string
  Default            string
  DateFormat         string
  DateFormatLocation *time.Location
  DateParseLayout    string
  DateParseLocation  *time.Location
}
```

## Contributing

We welcome contributions! Please see our [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute to this project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.