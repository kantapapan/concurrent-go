package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
	//"sync"
)

func main() {

	items := []string{"./ken_all1.csv", "./ken_all2.csv", "./ken_all3.csv"}

	// シングルスレッドの場合
	/*
		for _, item := range items {
			readCsv(item)
			fmt.Println(item + " : finished")
		}
	*/

	// goroutineの場合
	var wg sync.WaitGroup

	for _, item := range items {
		wg.Add(1)
		go func(item2 string) {
			defer wg.Done()
			readCsv(item2)
		}(item)

	}
	wg.Wait()

}

// FileLine ...
type FileLine struct {
	col1  string
	col2  string
	col3  string
	col4  string
	col5  string
	col6  string
	col7  string
	col8  string
	col9  string
	col10 string
}

// readCsv
func readCsv(filePath string) {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// CSVのReaderを用意
	r := csv.NewReader(file)

	// デリミタ(TSVなら\t, CSVなら,)設定
	r.Comma = ','

	// コメント設定(なんとコメント文字を指定できる!)
	r.Comment = '#'

	// 全部読みだす
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var count int
	var f FileLine
	// 各行でループ
	for _, v := range records {

		f = FileLine{
			col1:  v[0],
			col2:  v[4],
			col3:  v[5],
			col4:  v[5],
			col5:  v[5],
			col6:  v[5],
			col7:  v[5],
			col8:  v[5],
			col9:  v[5],
			col10: v[5],
		}
		count++
	}
	fmt.Println(count)
	fmt.Println(f)

}
