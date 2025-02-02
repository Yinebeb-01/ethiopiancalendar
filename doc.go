// Package ethiocal provides utilities for converting dates between the Ethiopian
// and Gregorian calendars and fetching important Ethiopian religious dates.
//
// This package offers a collection of APIs to facilitate:
// - Date conversion between Ethiopian and Gregorian calendars.
// - Retrieval of important Ethiopian religious festival dates.
//
// Available APIs:
//
// 1. BahireHasab API:
//   - Provides information about Ethiopian fasting and festival dates.
//
// 2. DateConverter APIs:
//   - Convert Gregorian dates to Ethiopian dates.
//   - Convert Ethiopian dates to Gregorian dates.
//
// Usage:
// This tool can be used in two ways:
//
// 1. Run on Your Own Server:
//   - Run the server by configuring it locally or on a cloud provider.
//
// 2. CLI Usage:
//   - A command-line interface (CLI) is available for direct usage.
//   - The CLI supports the following commands:
//   - `bahir`: Get Ethiopian fasting and religious festival dates for a given year.
//   - `convert`: Convert dates between Ethiopian and Gregorian calendars.
//   - `help`: Help on usage.
//
// Examples:
//
// To get bahire-hasab calendar for a specific ethiopian year:
//
//	ethiocal bahir 2017
//
// To convert a Gregorian date to Ethiopian:
//
//	ethiocal convert gtoe 2025-02-02
//
// To convert an Ethiopian date to Gregorian:
//
//	ethiocal convert etog 2017-05-25
//
// To enable server mode:
//
//	ethiocal --cli=false
//
// To integrate into your project, simply import the package and use the exported methods.
package main
