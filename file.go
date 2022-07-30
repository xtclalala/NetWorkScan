package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"strconv"
	"sync"
)

type worker struct {
	ip       string
	port     string
	user     string
	password string
}

var (
	inFilename  = global.File.InFileName
	outFilename = global.File.OutFileName
	sheet       = global.File.Sheet
	ip          = global.File.Ip
	port        = global.File.Port
	user        = global.File.User
	password    = global.File.Password
)

func ReadFile(workers *[]*worker) error {
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
	file := excelize.NewFile()
	defer file.Close()
	index := file.NewSheet(sheet)
	file.SetActiveSheet(index)

	i := 1
	data.Range(func(key, value any) bool {
		// 设置单元格的值
		si := strconv.Itoa(i)
		va := value.([]string)
		file.SetCellStr(sheet, "A"+si, fmt.Sprintf("%s", key))
		file.SetCellStr(sheet, "B"+si, va[0])
		file.SetCellStr(sheet, "C"+si, va[1])
		file.SetCellStr(sheet, "D"+si, va[2])
		file.SetCellStr(sheet, "E"+si, va[3])
		file.SetCellStr(sheet, "E"+si, va[4])
		i++
		return true
	})

	if err := file.SaveAs(outFilename); err != nil {
		return errors.Wrap(err, fmt.Sprintf("write %s is fail!", outFilename))
	}
	return nil
}
