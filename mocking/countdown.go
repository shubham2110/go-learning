package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdowntimer = 3
const finalWord = "Go!"

func Countdown(out io.Writer) {

	for i := countdowntimer; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
