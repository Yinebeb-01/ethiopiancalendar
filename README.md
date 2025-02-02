<p align="center">
<img src="logo.png" alt="logo" width="110" height="110">
</p>
<h1 align="center"><a href="https://pkg.go.dev/github.com/yinebebt/ethiocal">Ethiopian Calendar (ባሕረ-ሐሳብ)</a></h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/yinebebt/ethiocal.svg)](https://pkg.go.dev/github.com/yinebebt/ethiocal@v0.2.5)
[![ci-badge](https://github.com/yinebebt/ethiocal/actions/workflows/ci.yml/badge.svg)](https://github.com/yinebebt/ethiocal/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/yinebebt/ethiocal)](https://goreportcard.com/report/github.com/yinebebt/ethiocal)

## Description
The Ethiopian calendar(ባሕረ-ሀሳብ) is used to get Fasting and Holiday's specific date with in a year based on 
[EOTC](https://www.ethiopianorthodox.org/) calendar. It also designed to facilitate the conversion between Ethiopian dates (in the format yy-mm-dd) and 
Gregorian dates. Ethiopia follows its own calendar system, which consists of 13 months, each with 30 days. 

### Functionality
This tool allows users to:
* Get Ethiopian fasting and religious festival dates for a specific year.
* Convert Ethiopian dates to Gregorian dates.
* Convert Gregorian dates to Ethiopian dates.

### Usage
The tool can be used in two ways:

1. **Run on Your Own Server:**
    - Run the server by configuring it locally or on a cloud provider.

2. **CLI Usage:**
    - A comprehensive command-line interface (CLI) is available for direct usage.
    - The CLI supports the following commands:
        - `bahir`: Get Ethiopian fasting and religious festival dates for a given year.
        - `convert`: Convert dates between Ethiopian and Gregorian calendars.
        - `help`: Help on usage.

#### Examples:

To get religious dates for a specific Ethiopian year:
```bash
ethiocal bahir 2017
```

To convert a Gregorian date to Ethiopian:
```bash
ethiocal convert gtoe 2025-02-02
```

To convert an Ethiopian date to Gregorian:
```bash
ethiocal convert etog 2017-05-25
```

To enable server mode:
```bash
ethiocal --cli=false
```

### Installation
Install using the following Go command:
```bash
go install github.com/yinebebt/ethiocal@latest
```
