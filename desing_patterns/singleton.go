package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (d Database) CreateSingleConnection() {
	fmt.Println("Creating single database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

var db *Database
var lock sync.Mutex

func getDatabaseIntance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creando conexión a DB")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseIntance()
		}()
	}
	wg.Wait()
}
