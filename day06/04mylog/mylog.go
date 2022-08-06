package main

import (
	"fmt"
	"log"
	"os"
	"time"
)
const (
	Debug = iota
	Trace
	Info
	Warning
	Error 
	Fatal 
)
type human struct{
	name string
	age int
}
func (h human)Showname(){
	fmt.Println(h.name)	
}


func main(){
	/*
	// 直接将日志输出到标准输出
	for {
		log.Println("This is a log msg")
		time.Sleep(3*time.Second)
		nowTime := time.Now()
		fmt.Println(nowTime.Format("2006/01/02 15:04:05"))
	}
	*/
	/*

	// 将日志输出到指定的 buf bytes.Buffer
	var buf bytes.Buffer
	logger := log.New(&buf,"logger: ",log.Lshortfile)
	logger.Print("Hello, log file!")
	fmt.Print(&buf)
	*/
	

	
	// 将日志输出到指定的文件中
	file,err := os.OpenFile("xxx.log", os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Println(err)
		return
	}
	nowTime := time.Now()
	logger := log.New(file, nowTime.Format("2006/01/02 15:04:05 ") , log.Lshortfile)
	logger.SetFlags(8)
	logger.SetPrefix(nowTime.Format("2006-01-02 15:04:05 "))
	logger.Output(1, "Logger: This is a log file")
//	logger.f("Logger: This is a xxx log")
//	logger.Print("This is a xxx log")
	Tom := human{"Tom", 18}
	Tom.Showname()
	fmt.Printf("%s:%d\n", "Debug",Debug)
	fmt.Printf("%s:%d\n", "Info",Info)
	fmt.Printf("%s:%d\n", "Trace",Trace)
	fmt.Printf("%s:%d\n", "Warning",Warning)
	fmt.Printf("%s:%d\n", "Error",Error)
	fmt.Printf("%s:%d\n", "Fatal", Fatal)

}
