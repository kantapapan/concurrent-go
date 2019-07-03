package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 1度に取得するbulk行
	bulkCount := 2

	fr, err := os.Open(`ken_all1.csv`)
	failOnError(err)
	defer fr.Close()

	reader := csv.NewReader(fr)

	for {
		lines := make([][]string, 0, bulkCount)
		isLast := false
		for i := 0; i < bulkCount; i++ {
			line, err := reader.Read()
			if err == io.EOF {
				isLast = true
				break
			} else if err != nil {
				panic(err)
			}
			lines = append(lines, line)
		}
		//fmt.Println(lines)

		//bulkdata 単位で何かしらの処理
		execBulkLines(lines)

		if isLast {
			fmt.Println("break")
			break
		}
	}
}

// FileLine ...
type FileLine struct {
	col0 string
	col1 string
	col2 string
	col3 string
	col4 string
}

func execBulkLines(lines [][]string) FileLine {

	var f FileLine
	// 各行でループ
	for _, line := range lines {
		//fmt.Println(line) // [ a b c]
		f = FileLine{
			col0: line[0],
			col1: line[1],
			col2: line[2],
			col3: line[3],
			col4: line[4],
		}
	}
	return f
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
