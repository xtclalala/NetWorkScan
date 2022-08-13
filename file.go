package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"io"
	"strings"
	"sync"
)

type worker struct {
	ip       string
	port     string
	user     string
	password string
}

func ReadFile(workers *[]*worker) error {
	var (
		inFilename = global.File.InFileName
		sheet      = global.File.Sheet
		ip         = global.File.Ip
		port       = global.File.Port
		user       = global.File.User
		password   = global.File.Password
	)
	file, err := excelize.OpenFile(inFilename)
	defer file.Close()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("read file %s is fail!", inFilename))
	}

	// 获取 Sheet1 上所有单元格
	rows, err := file.GetRows(sheet)
	for _, row := range rows {
		i := &worker{
			ip:       row[ip],
			port:     row[port],
			user:     row[user],
			password: row[password],
		}
		*workers = append(*workers, i)

	}
	return nil
}

func WriteFile(data *sync.Map) error {
	//var (
	//	sheet       = global.File.Sheet
	//	outFilename = global.File.OutFileName
	//)
	//file := excelize.NewFile()
	//defer file.Close()
	//index := file.NewSheet(sheet)
	//file.SetActiveSheet(index)
	//i := 1
	data.Range(func(key, value any) bool {
		// 设置单元格的值
		//values := value.(string)
		//column := string(65) + strconv.Itoa(i)
		//file.SetCellStr(sheet, column, strings.Trim(values, "\n"))
		//i++
		//
		//return true
		res := strings.Split(strings.ReplaceAll(value.(string), "nohup.out", "*"), "||")
		for i, re := range res {
			dec := json.NewDecoder(strings.NewReader(re))
			var parse Temp
			for {
				if err := dec.Decode(&parse); err == io.EOF {
					break
				} else if err != nil {
					break
				}
			}
			fmt.Printf("index: %v, value: %#v\n", i, parse.Value)
		}
		return true
	})
	return nil
	//if err := file.SaveAs(outFilename); err != nil {
	//	return errors.Wrap(err, fmt.Sprintf("write %s is fail!", outFilename))
	//}
	//return nil
}

type Temp struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
