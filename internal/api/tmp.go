package api

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
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
