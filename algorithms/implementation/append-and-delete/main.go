package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
	if len(s)+len(t)-int(k) < 0 {
		return "Yes"
	}

	var short, long []byte

	if len(s) > len(t) {
		short, long = []byte(t), []byte(s)
	} else {
		short, long = []byte(s), []byte(t)
	}

	matched := func() int {
		for i := range short {
			if short[i] != long[i] {
				return i
			}
		}

		return len(short)
	}()

	ops := len(s) + len(t) - matched*2
	switch {
	default:
		return "No"
	case ops > int(k):
		return "No"
	case ops%2 == int(k%2):
		return "Yes"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
