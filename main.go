package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	// read config file
	InitConfig()

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
				//out, _ := s.RunCmd(cmd)
				res := s.ScanOS()
				values = append(values, res...)
				data.Store(worker.ip, values)
			}

		}(item)
		fns = append(fns, fn)

	}
	Start(fns)
	data.Range(func(key, value any) bool {
		fmt.Printf("key:%v value: %v\n", key, value)
		return true
	})
	if err := WriteFile(data); err != nil {
		log.Fatalf(err.Error())
	}

}
