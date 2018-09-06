package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	//fmt.Fprint(out, "3")

	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "GO!")
}

func main() {
	Countdown(os.Stdout)
}
