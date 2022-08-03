package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var cliConfigPath = flag.String("path", "../config.yml", "config fail path")

func main() {
	flag.Parse()

	// read config file
	InitConfig(*cliConfigPath)

	// read check file
	workers := new([]*worker)
	if err := ReadFile(workers); err != nil {
		log.Fatalf(err.Error())
	}

	data := &sync.Map{}
	var fns []func()
	for _, item := range *workers {
		var fn = func(worker *worker) func() {
			return func() {
				// connect
				s := NewSsh(worker.ip, worker.port, worker.user, worker.password)
				s.Connect()
				s.GetOS()
				values := s.Save()
				// you can do something, run diy cmd
				res := s.ScanOS()
				values = append(values, res...)
				data.Store(worker.ip, values)
			}

		}(item)
		fns = append(fns, fn)

	}
	Start(fns)
	if err := WriteFile(data); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("按任意键退出...")
	in := make([]byte, 1)
	os.Stdin.Read(in)

}
