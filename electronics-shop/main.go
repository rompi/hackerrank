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
			if(ar[i] < ar[j]) {
				ar[i], ar[j] = swap(ar[i],ar[j])
			}
		}
	}
	return ar
}

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	keyboards = sort(keyboards)
	drives = sort(drives)
	var res []int32
	for k := 0; k < len(keyboards); k++ {
		for d := 0; d < len(drives); d++ {
			if (keyboards[k] + drives[d]) <= b {
				res = append(res, keyboards[k] + drives[d])
				break
			}
		}
	}
	if len(res) > 0 {
		lastResult := sort(res)[0]
		return lastResult
	}

	return -1
}

// https://www.hackerrank.com/challenges/electronics-shop
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	bnm := strings.Split(readLine(reader), " ")

	bTemp, err := strconv.ParseInt(bnm[0], 10, 64)
	checkError(err)
	b := int32(bTemp)

	nTemp, err := strconv.ParseInt(bnm[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(bnm[2], 10, 64)
	checkError(err)
	m := int32(mTemp)

	keyboardsTemp := strings.Split(readLine(reader), " ")

	var keyboards []int32

	for keyboardsItr := 0; keyboardsItr < int(n); keyboardsItr++ {
		keyboardsItemTemp, err := strconv.ParseInt(keyboardsTemp[keyboardsItr], 10, 64)
		checkError(err)
		keyboardsItem := int32(keyboardsItemTemp)
		keyboards = append(keyboards, keyboardsItem)
	}

	drivesTemp := strings.Split(readLine(reader), " ")

	var drives []int32

	for drivesItr := 0; drivesItr < int(m); drivesItr++ {
		drivesItemTemp, err := strconv.ParseInt(drivesTemp[drivesItr], 10, 64)
		checkError(err)
		drivesItem := int32(drivesItemTemp)
		drives = append(drives, drivesItem)
	}

	/*
	 * The maximum amount of money she can spend on a keyboard and USB drive, or -1 if she can't purchase both items
	 */

	moneySpent := getMoneySpent(keyboards, drives, b)

	fmt.Fprintf(writer, "%d\n", moneySpent)

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
