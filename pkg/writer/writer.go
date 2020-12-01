package writer

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Writer is a interface that write something to somewhere.
type Writer interface {
	Write([]string) error
	Flush() error
}

// TSVWriter is a struct that write tab separate table to stdout.
type TSVWriter struct {
	w *tabwriter.Writer
}

// NewTSVWriter is a constructor that make new TSVWriter.
func NewTSVWriter() *TSVWriter {
	return &TSVWriter{
		w: tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0),
	}
}

// Write is a function that write string slice to buffer.
func (w *TSVWriter) Write(record []string) error {
	str := strings.Join(record[:], "\t")
	if _, err := fmt.Fprintln(w.w, str); err != nil {
		return err
	}

	return nil
}

// Flush is a function that write buffet to stdout.
func (w *TSVWriter) Flush() error {
	if err := w.w.Flush(); err != nil {
		return err
	}

	return nil
}
