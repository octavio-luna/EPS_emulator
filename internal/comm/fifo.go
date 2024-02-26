package comm

import (
	"fmt"
	"os"
)

type Writer struct {
	writeFifo *os.File
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.writeFifo.Write(p)
}

func (w *Writer) Close() error {
	return w.writeFifo.Close()
}

func New() (readFifo *os.File, writer *Writer, err error) {
	// Open FIFOs for reading and writing
	fmt.Println("Opening FIFOs...")
	readFifo, err = os.OpenFile("/tmp/eps_read_fifo", os.O_RDWR, os.ModeNamedPipe) // We have to open the FIFO in read-write mode because read/writeOnly are blocking until the other end is opened
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Read FIFO opened successfully")

	writeFifo, err := os.OpenFile("/tmp/eps_write_fifo", os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		readFifo.Close()
		return nil, nil, err
	}
	fmt.Println("Write FIFO opened successfully")

	return readFifo, &Writer{writeFifo}, nil
}
