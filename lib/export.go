package lib

import (
	"bufio"
	"os"
)

// FileExporter holds buffered file writing capabilities
type FileExporter struct {
	Path   string
	File   *os.File
	Writer *bufio.Writer
}

// NewFileExporter validate path and build new FileExporter instance
func NewFileExporter(path string) (FileExporter, error) {
	file, err := os.Create(path)
	if err != nil {
		return FileExporter{}, err
	}
	return FileExporter{
		Path:   path,
		File:   file,
		Writer: bufio.NewWriter(file),
	}, nil
} // END NewFileExporter

// WriteString exposes FileExporter as Writer
func (fe *FileExporter) WriteString(s string) (int, error) {
	return fe.Writer.WriteString(s)
} // END WriteString

// Write exposes FileExporter as Writer
func (fe *FileExporter) Write(p []byte) (int, error) {
	return fe.Writer.Write(p)
} // END Write

// Close exposes FileExporter as Writer
func (fe *FileExporter) Close() error {
	err := fe.Writer.Flush()
	fe.File.Close()
	// what about the potential close error?
	return err
} // END Close
