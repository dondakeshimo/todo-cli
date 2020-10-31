package writer

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Writer interface {
	Write([]string) error
	Flush()
}

type TSVWriter struct {
	w *tabwriter.Writer
}

func NewTSVWriter() *TSVWriter {
	return &TSVWriter{
		w: tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0),
	}
}

func (w *TSVWriter) Write(record []string) error {
	str := strings.Join(record[:], "\t")
	if _, err := fmt.Fprintln(w.w, str); err != nil {
		return err
	}

	return nil
}

func (w *TSVWriter) Flush() {
	w.w.Flush()
}
