package main

/*
import (
	"fmt"
	"os"
)
func fileWrite(){
	f,err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil{
		fmt.Println("Open file failed: ", err)
	}
	defer f.Close()
	f.Write([]byte("hello world\n"))
	f.WriteString("I am learning Go!")
}

func main(){
	fileWrite()
}
*/
/*
import (
	"fmt"
	"bufio"
	"os"
)

func fileWriteWithbufio(){
	f,err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Println("file open failed: ",err)
		return
	}
	defer f.Close()
	// 创建一个缓存
	writer := bufio.NewWriter(f)
	// 将字符串数据写入缓存
	writer.WriteString("hello world\n")
	// 将缓存数据刷入文件
	writer.Flush()
}
func main(){
	fileWriteWithbufio()
}
*/
import (
	"fmt"
	"io/ioutil"
)

func fileWriteWithioutil(){
	str := "hello world!"
	err := ioutil.WriteFile("test.txt", []byte(str), 0644)
	if err != nil{
		fmt.Println("file open failed: ", err)
		return
	}

}
func main(){
	fileWriteWithioutil()
}