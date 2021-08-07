package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	channel := make(chan int,10)
	for i := 0;i < 10;i++{
		channel <- i
	}
	close(channel)
	for i := 0; i < 10;i++{
		go func() {
			defer wg.Done()
			fmt.Println(<-channel)
		}()
	}
	wg.Wait()




}

type T_Table struct {
	ID int64 `json:"id" gorm:"id"`
	A int    `json:"a" gorm:"a"`
	B int    `json:"b" gorm:"b"`
	C int    `json:"c" gorm:"c"`
}

type WhereTTable struct {
	ID *int64 `operater:""`
	IDGt *int64 `operater:">"`
	A *int    `json:"a" gorm:"a"`
	B *int    `json:"b" gorm:"b"`
	C *int    `json:"c" gorm:"c"`
}

type Opt struct {
	limit *int32
	offset *int32
}
func find(c context.Context,whereTTable *WhereTTable,opt *Opt)[]*T_Table{
	po := make([]*T_Table,0)
	err := gormx.find(c,whereTTable,po,
	return po
}

find()