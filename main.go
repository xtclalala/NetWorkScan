package main

import (
	"log"
	"sync"
)

func main() {
	// 读配置文件
	InitConfig()

	// 读需要检查的文件
	workers := *(new([]*worker))
	if err := ReadFile(&workers); err != nil {
		log.Fatalf(err.Error())
	}
	// 获取cmd
	cmd := "cd / & ls"
	w := &sync.WaitGroup{}
	data := &sync.Map{}

	for _, item := range workers {
		worker := item
		w.Add(1)
		go func() {
			// 连接
			defer w.Done()
			s := NewSsh(worker.ip, worker.port, worker.user, worker.password)
			s.Connect()
			out, _ := s.RunCmd(cmd)
			// do something
			data.Store(worker.ip, []string{worker.port, worker.user, worker.password, cmd, out})

		}()
	}
	// 获取返回的数据进行处理
	w.Wait()

	if err := WriteFile(data); err != nil {
		log.Fatalf(err.Error())
	}

}

// 1.批量读取文件中的 ip 端口 用户 密码 	finish
// 2.连接 								finish
// 3.识别linux操作系统的发行版本
// 4.save     							finish
// 5.读配置文件							finish
