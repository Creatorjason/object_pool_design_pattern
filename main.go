package main

import (
	"fmt"
	"log"
	"strconv"
)

func main(){
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3 ; i++{
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := InitPool(connections)
	if err != nil{
		log.Fatalf("Init pool error: %s", err)
	}

	conn1, err := pool.loan()
	if err != nil{
		log.Fatalf("Pool Loan error: %s", err)
	}
	conn2, err := pool.loan()
	if err != nil{
		log.Fatalf("Pool Loan error: %s", err)
	}
	_, err = pool.loan()
if err != nil{
	log.Fatalf("Pool Loan error: %s", err)
}
// _, err = pool.loan()
// if err != nil{
	// log.Fatalf("Pool Loan error: %s", err)
// }
	pool.receive(conn1)
	pool.receive(conn2)
	fmt.Println(pool)
}