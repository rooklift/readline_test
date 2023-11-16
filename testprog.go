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

	output(n)
}

func output(n int) {

	arr := make([]byte, n + 1)

	for i := 0; i < n; i++ {
		if i == n - 1 {
			arr[i] = 'E'			// E for END - the last character except the newline
		} else {
			arr[i] = 'c'			// c for character
		}
	}
	arr[n] = '\n'

	os.Stdout.Write(arr[:])

}
