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

func New(read, write string) (readFifo *os.File, writer *Writer, err error) {
	// Open FIFOs for reading and writing
	read_fifo_path := fmt.Sprintf("/tmp/%s", read)
	write_fifo_path := fmt.Sprintf("/tmp/%s", write)
	fmt.Println("Opening FIFOs... ", read_fifo_path, write_fifo_path)

	readFifo, err = os.OpenFile(read_fifo_path, os.O_RDWR, os.ModeNamedPipe) // We have to open the FIFO in read-write mode because read/writeOnly are blocking until the other end is opened
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Read FIFO opened successfully")

	writeFifo, err := os.OpenFile(write_fifo_path, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		readFifo.Close()
		return nil, nil, err
	}
	fmt.Println("Write FIFO opened successfully")

	return readFifo, &Writer{writeFifo}, nil
}
