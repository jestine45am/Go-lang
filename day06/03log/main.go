package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main(){
	f1,err := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("Open log file failed, ", err)
		return
	}
	log.SetOutput(f1)
	for {
		log.Println("这是一条日志消息！")
		time.Sleep(5 * time.Second)
	}
}