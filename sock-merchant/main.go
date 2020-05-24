package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func swap(a, b int32) (int32, int32) {
	return b, a
}

func sort(ar []int32) []int32 {
	n := len(ar)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if(ar[i] > ar[j]) {
				ar[i], ar[j] = swap(ar[i],ar[j])
			}
		}
	}
	return ar
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	var counter int32
	counter = 0
	arr := sort(ar)
	for i := 0; i < len(arr); i++ {
		if i+1 == len(arr) {
			break
		} else if arr[i] == arr[i+1] {
			counter++
			i += 1
		}
	}
	return counter
}

// https://www.hackerrank.com/challenges/sock-merchant
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

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
