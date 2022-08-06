package main

import (
	"bufio"
	"fmt"
	"io"
	//"io/ioutil"
	"os"
)

/*
func readFileWithSlice(){
	// 打开文件
	f1,err := os.Open("./main.go")
	if err != nil{
		fmt.Println("Open file failed", err)
		return
	}
	// 关闭文件
	defer f1.Close()
	for{
		var tmp = make([]byte, 128)
		// 读取文件
		n,err := f1.Read(tmp)
		if err != nil{
			fmt.Println("Read file failed", err)
			return
		}

		fmt.Println(string(tmp[:n]))
		// 如果最后读取的字节小于 128 说明已经全部读取完毕
		if n < len(tmp) {
			return
		}
	}
}
*/
/*
func readFileWithIo(){
	f1,err := os.Open("./main.go")
	if err != nil{
		fmt.Printf("File open failed, %v\n", err)
		return
	}
	defer f1.Close()
	var buf = [128]byte{}
	for{
		n,error := f1.Read(buf[:])
		if error == io.EOF{
			fmt.Println("File read completed")
			return
		}
		if error != nil{
			fmt.Printf("File read failed, %v\n", error)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
*/
func readFileWithBufio(){
	f1,err := os.Open("./main.go")
	if err != nil{
		fmt.Printf("Open file failed, %v\n", err)
		return
	}
	defer f1.Close()
	reader := bufio.NewReader(f1)
	for{
		str,error := reader.ReadString('\n')
		if error == io.EOF{
			fmt.Println("Read file completed!")
			return
		}
		if error != nil{
			fmt.Printf("Read file failed, %v\n",error)
			return
		}
		fmt.Print(str)
	}
}

/*
func readFileWithiountil(){
	buf,err := ioutil.ReadFile("./main.go")
	if err != nil{
		fmt.Println("Read file failed!", err)
		return
	}
	fmt.Println(string(buf))
}
*/

func main(){
	//readFileWithIo()
	//readFileWithSlice()
	readFileWithBufio()
	//readFileWithiountil()
}
