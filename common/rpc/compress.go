package rpc

import (
	"bufio"
	"compress/flate"
	"io"
)

type compressor interface {
	flushWriter
	io.Reader
}

type flushWriter interface {
	io.Writer
	Flush() error
}

type Compressor struct {
	zr io.Reader
	zw flushWriter
}

type flushWraper struct {
	zw    flushWriter
	flush func() error
}

func (f *flushWraper) Flush() error {
	if err := f.zw.Flush(); err != nil {
		return err
	}
	return f.flush()
}
func (f *flushWraper) Write(p []byte) (n int, err error) {
	return f.zw.Write(p)
}

func newFlushWraper(w flushWriter, flush func() error) flushWriter {
	return &flushWraper{
		zw:    w,
		flush: flush,
	}
}

func newCompressor(r io.Reader, w io.Writer, compress string) (*Compressor, error) {
	var zr io.Reader
	var zw flushWriter
	var err error

	bw := bufio.NewWriterSize(w, writeBufferSize)
	compress = ""
	switch compress {
	case "deflate":
		zr = flate.NewReader(r)
		zw, err = flate.NewWriter(bw, flate.DefaultCompression)
		if err != nil {
			return nil, err
		}
		zw = newFlushWraper(zw, bw.Flush)
	default:
		br := bufio.NewReaderSize(r, readBufferSize)
		zr = br
		zw = bw
	}

	return &Compressor{
		zr: zr,
		zw: zw,
	}, nil
}

func (c *Compressor) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}
func (c *Compressor) Write(p []byte) (n int, err error) {
	return c.zw.Write(p)
}
func (c *Compressor) Flush() (err error) {
	return c.zw.Flush()
}
