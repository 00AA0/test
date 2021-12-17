package api

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"sort"
	"sync"
)

// https://www.wenjiangs.com/doc/canal-introduce   canal文档
// https://github.com/go-mysql-org/go-mysql  使用例子
type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Println("kaishi")
	fmt.Printf("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func StartCanal() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"
	cfg.Password = "123456789"
	cfg.Dump.ExecutionPath = "/usr/local/mysql/bin/mysqldump"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "test"
	cfg.Dump.Tables = []string{"tblUser"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
}

func Goruntine() {

	wg := &sync.WaitGroup{}

	ch1 := make(chan int)
	//ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	for j := 0; j < 3; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for i := j * 10; i < 10+j*10; i++ {
				ch1 <- i
			}
		}(j)
	}

	go func() {
		wg.Wait()
		close(ch1)
	}()
	// 在主goroutine中从ch2中接收值打印
	var arr []int
	for i := range ch1 { // 通道关闭后会退出for range循环
		arr = append(arr, i)
	}
	sort.Ints(arr)
	fmt.Println(arr)
}
