package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
func process(conn net.Conn){
	tmp := [128]byte{}
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err2 := conn.Read(tmp[:])
		if err2 != nil {
			fmt.Println("Read msg failed, err:", err2)
			return
		}
		fmt.Println(string(tmp[:n]))

		fmt.Print("请输入：")
		msg,_ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		_,err3 := conn.Write([]byte(msg))
		if err3 != nil {
			fmt.Println("Write msg failed, err: ", err3)
			return
		}
	}

}

func main(){
	// 开启本地服务
	listener,err := net.Listen("tcp","127.0.0.1:33333")
	if err != nil{
		fmt.Println("start server failed, err:", err)
		return
	}
	// 等待客户端的连接
	for {
		conn, err1 := listener.Accept()
		if err1 != nil{
			fmt.Println("Accept connect failed, err:", err1)
			return
		}
		// 与客户端建立通信
		go process(conn)
	}

}