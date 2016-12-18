package main

import (
	"bufio"
	"fmt"
	"github.com/dchest/cssmin"
	"io"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprintf(os.Stderr, "Not reading from a pipe, aborting.\n")
		os.Exit(1)
	} else {
		r := bufio.NewReader(os.Stdin)
		buf := make([]byte, 0, 4*1024)
		result := make([]byte, 0, 1024*1024)
		for {
			n, err := r.Read(buf[:cap(buf)])
			buf = buf[:n]
			if n == 0 {
				if err == nil {
					continue
				}
				if err == io.EOF {
					break
				}
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			result = append(result, buf...)
			if err != nil && err != io.EOF {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		}
		fmt.Printf("%s\n", cssmin.Minify(result))
	}
}
