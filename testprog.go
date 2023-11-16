package main

import (
	"os"
	"strconv"
)

const DEFAULT_N = 256

func main() {

	var n int
	var err error

	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			n = DEFAULT_N
		}
	} else {
		n = DEFAULT_N
	}

	for i := 0; i <= n; i += 100000 {
		output(i)
	}
}

func output(n int) {

	arr := make([]byte, n + 1)

	for i := 0; i < len(arr); i++ {
		arr[i] = '.'
	}
	arr[len(arr) - 1] = '\n'

	if len(arr) > 3 {
		copy(arr[len(arr) - 4 : len(arr) - 1], "END")
	}

	os.Stdout.Write(arr)

}
