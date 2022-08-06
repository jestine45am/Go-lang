package main

import (
	"fmt"
	"time"
)
/*
func main(){
	now := time.Now()
	fmt.Println(now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(time.Unix(now.Unix(),0))

	fmt.Println(time.Second)
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
}
*/
func main(){
	now := time.Now()
	fmt.Println(now.Location())
	later := now.Add(1 * time.Hour)
	fmt.Println("1 hour later time: ", later)
/*
	// 定时器
	timer := time.Tick(1 * time.Second)
	for t := range timer{
		fmt.Println(t)
	}
	*/
	
	/*
	//通过指定形式的时间格式，返回字符串形式的时间格式
	timer := time.Now().Format("03:04:05 01/02/2006")
	fmt.Println(timer)
	timer1 := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Println(timer1)
	timer2, err := time.Parse("2006-01-02", "2022-04-24")
	if err != nil{
		fmt.Println("time error: ",err)
		return
	}
	fmt.Println(timer2)
	*/
//	lc,_ := time.LoadLocation("Local")
	timer01,_:= time.ParseInLocation("2006-01-02 15:04:05", "2022-04-26 00:00:00", time.Local)
	fmt.Println(timer01)
	fmt.Println(now.Sub(timer01))
	fmt.Println(timer01.Sub(now))
	time.Sleep(time.Duration(100) * time.Second)
}