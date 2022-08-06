package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main(){
	rand.Seed(time.Now().UnixMicro())	
	for i := 0; i < 10; i++{
		fmt.Println(rand.Int(), rand.Intn(10))
	}
}