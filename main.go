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
	// 获取cmd
	cmd := "cd / & ls"
	data := &sync.Map{}
	var fns []func()
	for _, item := range *workers {
		worker := item
		var fn = func() {
			// 连接
			s := NewSsh(worker.ip, worker.port, worker.user, worker.password)
			s.Connect()
			// you can do something
			out, _ := s.RunCmd(cmd)
			data.Store(worker.ip, []string{worker.ip, worker.port, worker.user, worker.password, cmd, out})
		}
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

// 1.识别linux操作系统的发行版本    cat /etc/os-release openEuler centos PRETTY_NAME hostnamectl
// 2.命令
// 3.超时处理
