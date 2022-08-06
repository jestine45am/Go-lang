package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 与服务端建立连接
	conn, err := net.Dial("tcp","127.0.0.1:33333")
	if err != nil{
		fmt.Println("Dial server failed, err:",err)
		return
	}
	// 发送消息
	reader := bufio.NewReader(os.Stdin)
	tmp := [128]byte{}
	for {
		fmt.Print("请输入：")
		msg,_ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit"{
			break
		}
		_,err1 := conn.Write([]byte(msg))
		if err1 != nil{
			fmt.Println("Write msg failed, err: ", err1)
			return
		}
		n,err2 := conn.Read(tmp[:])
		if err2 != nil {
			fmt.Println("Read msg failed, err: ", err2)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
	conn.Close()

}