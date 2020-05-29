package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sum(arr []int32) int32 {
	var sum int32
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func solve(arr []int32) string {
	for i := 0; i < len(arr); i++ {
		if sum(arr[:i]) == sum(arr[i:]){
			return "YES"
		}
	}
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)


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

	result := solve(ar)
	fmt.Println(result)
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
