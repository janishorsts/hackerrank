package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the permutationEquation function below.
func permutationEquation(p []int32) []int32 {
	r := make([]int32, len(p))

	f := func(v int32) int32 {
		for i := range p {
			if p[i] == v {
				return int32(i + 1)
			}
		}
		return 0
	}

	for i := range p {
		r[i] = f(f(int32(i + 1)))
	}

	return r
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

	pTemp := strings.Split(readLine(reader), " ")

	var p []int32

	for i := 0; i < int(n); i++ {
		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		pItem := int32(pItemTemp)
		p = append(p, pItem)
	}

	result := permutationEquation(p)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
