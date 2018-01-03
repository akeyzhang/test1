package main

import (
	"sync"
	"test1/cmd/pool"
	"sync/atomic"
	"log"
	"time"
	"math/rand"
	"io"
)


const (
	maxGoroutine = 5
	poolRes = 2
)
var idCounter int32

type dbConnection struct {
	ID int32
}


func (db *dbConnection) Close() error {
	log.Println("關畢連接:",db.ID)
	return nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
	p,err := pool.New(createConnection,poolRes)
	if err != nil {
		log.Println(err)
		return
	}
	for query:=0;query<maxGoroutine;query++ {
		go func(q int) {
			dbQuery(q,p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("開始關畢資源池")
	p.Close()
}


func dbQuery(query int,pool *pool.Pool)  {
	conn,err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	log.Printf("第%d個查詢,使用的是ID:%d的數據庫連接\n",query,conn.(*dbConnection).ID)
}

func createConnection() (io.Closer,error) {
	id:=atomic.AddInt32(&idCounter,1)
	return &dbConnection{id},nil
}

