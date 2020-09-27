package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("ispmpan02o.log-20180924.gz")
	defer file.Close()
	lc, _ := lineCounter(file)
	fmt.Println("number of lines:", lc)
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
