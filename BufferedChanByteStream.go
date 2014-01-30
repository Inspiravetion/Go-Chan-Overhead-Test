package main

import (
	"io"
	"os"
)

type ByteStream struct {
	c <-chan byte
}

func main() {
	for i := 0; i < 100; i++ {
		stream := New_Byte_Stream(os.Args[1])
		for _, err := stream.get_byte(); err == nil; _, err = stream.get_byte() {
		}
	}
}

func New_Byte_Stream(path string) *ByteStream {
	bs := new(ByteStream)
	bs.c = get_bytes(path)

	return bs
}

func (this *ByteStream) get_byte() (byte, error) {
	if b, ok := <-this.c; ok {
		return b, nil
	} else {
		return b, io.EOF
	}
}

func get_bytes(path string) <-chan byte {
	c := make(chan byte, 1024) //10, 25, 50, 1024

	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	go func() {
		buff := make([]byte, 1024)

		defer close(c)
		defer input.Close()

		for {
			_, err := input.Read(buff)
			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			for _, b := range buff {
				c <- b
			}
		}
	}()

	return c
}
