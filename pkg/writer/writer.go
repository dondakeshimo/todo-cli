package writer

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Writer is a interface that write something to somewhere.
type Writer interface {
	Header([]string) error
	Write([]string) error
	Flush() error
}

// TSVWriter is a struct that write tab separate table to stdout.
type TSVWriter struct {
	w *tablewriter.Table
}

// NewTSVWriter is a constructor that make new TSVWriter.
func NewTSVWriter() *TSVWriter {
	tsvw := &TSVWriter{
		w: tablewriter.NewWriter(os.Stdout),
	}
	tsvw.w.SetAutoFormatHeaders(false) // for preventing header from being made upper case automatically
	return tsvw
}

// Header is a function that write header string slice to buffer.
func (w *TSVWriter) Header(record []string) error {
	w.w.SetHeader(record)

	return nil
}

// Write is a function that write string slice to buffer.
func (w *TSVWriter) Write(record []string) error {
	w.w.Append(record)

	return nil
}

// Flush is a function that write buffet to stdout.
func (w *TSVWriter) Flush() error {
	w.w.Render()

	return nil
}
