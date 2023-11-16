package main

import (
	"os"
	"strconv"
)

func main() {

	var n int

	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
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
