package main

import (
	"io"
	"os"
	"runtime"
)

type ByteStream struct {
	stream []byte
	c      <-chan []byte
	index  int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < 100; i++ {
		stream := New_Byte_Stream(os.Args[1])
		for _, err := stream.get_byte(); err == nil; _, err = stream.get_byte() {
		}
	}
}

func New_Byte_Stream(path string) *ByteStream {
	bs := new(ByteStream)
	bs.c = get_bytes(path)
	bs.index = -1

	return bs
}

func (this *ByteStream) get_byte() (byte, error) {
	if this.index == -1 {
		this.stream = <-this.c
		this.index = 0
	}

	var b byte
	this.index++

	if this.index == cap(this.stream) {
		b = this.stream[this.index-1]

		if cap(this.stream) == len(this.stream) {
			var ok bool
			if this.stream, ok = <-this.c; ok {
				this.index = 0
				return b, nil
			}

			return b, io.EOF
		}

		return b, io.EOF
	}

	return this.stream[this.index-1], nil
}

func get_bytes(path string) <-chan []byte {
	c := make(chan []byte)

	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	go func() {

		defer close(c)
		defer input.Close()

		for {
			buff := make([]byte, 1024)

			_, err := input.Read(buff)
			if err == io.EOF {
				return
			}

			if err != nil {
				panic(err)
			}

			c <- buff
		}
	}()

	return c
}
