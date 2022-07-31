package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	test()

	// 读配置文件
	//InitConfig()
	//
	//// 读需要检查的文件
	//workers := *(new([]*worker))
	//if err := ReadFile(&workers); err != nil {
	//	log.Fatalf(err.Error())
	//}
	//// 获取cmd
	//cmd := "cd / & ls"
	//w := &sync.WaitGroup{}
	//data := &sync.Map{}
	//
	//for _, item := range workers {
	//	worker := item
	//	w.Add(1)
	//	go func() {
	//		// 连接
	//		defer w.Done()
	//		s := NewSsh(worker.ip, worker.port, worker.user, worker.password)
	//		s.Connect()
	//		out, _ := s.RunCmd(cmd)
	//		// do something
	//		data.Store(worker.ip, []string{worker.port, worker.user, worker.password, cmd, out})
	//
	//	}()
	//}
	//// 获取返回的数据进行处理
	//w.Wait()
	//
	//if err := WriteFile(data); err != nil {
	//	log.Fatalf(err.Error())
	//}

}

// 1.识别linux操作系统的发行版本    cat /etc/os-release openEuler centos PRETTY_NAME
// 2.并发限速
// 3.超时处理  					finish

func test() {
	workers := []string{"11", "33", "22", "99", "66", "44"}
	fns := make([]func(ctx *context.Context), 0)
	for range workers {
		fn := NewWorker(doSomething)
		fns = append(fns, fn)
	}
	Start(fns)
}

// 干活
func doSomething(c context.Context, cancel context.CancelFunc) {
	for {
		select {
		default:
			time.Sleep(3 * time.Second)
			fmt.Printf("ssh: %s\n", "1")
			cancel()
		case <-c.Done():
			return

		}
	}

}
