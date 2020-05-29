package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	n := len(arr)
	var totalPositive, totalNegative, totalZero int32
	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			totalNegative += 1
		} else if arr[i] > 0 {
			totalPositive += 1
		} else {
			totalZero += 1
		}
	}

	fmt.Println(float32(totalPositive)/float32(n))
	fmt.Println(float32(totalNegative)/float32(n))
	fmt.Println(float32(totalZero)/float32(n))
}

// https://www.hackerrank.com/challenges/plus-minus
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
