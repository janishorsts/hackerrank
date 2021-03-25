package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const maxDays = 50

var likes []int32

func init() {
	likes = make([]int32, maxDays+1)

	var shared int32 = 5
	var liked int32

	for i := 1; i <= maxDays; i++ {
		likedOnDay := shared / 2

		liked += likedOnDay
		shared = likedOnDay * 3

		likes[i] = liked
	}
}

// Complete the viralAdvertising function below.
func viralAdvertising(n int32) int32 {
	return likes[n]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := viralAdvertising(n)

	fmt.Fprintf(writer, "%d\n", result)

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
