package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)


func main(){
	// 打开源文件获取文件句柄
	f1,err := os.Open("./test.txt")
	if err != nil{
		fmt.Println("Open file failed: ", err)
		return 
	}
	// 创建一个临时文件
	f2,err := os.OpenFile("./tmp.txt",os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Println("Open file failed: ", err)
		return 
	}
	reader := bufio.NewReader(f1)
	writer := bufio.NewWriter(f2)
	// 指定需要新插入数据在第 n 行，插入的字符串是 str
	str := "hello world\n"
	i := 1
	n := 5
	// 循环地从源文件的缓冲区中读取每一行，然后写入到临时文件的缓冲区中
	// 如果读完了源文件的前 n 行数据，那么将字符串 str 插入到临时文件的缓冲区中
	// 继续读取源文件缓冲区中的每一行，然后写入到临时文件的缓冲区中
	// 如果源文件的缓冲区中的数据读取完毕，写入到临时文件的缓冲区中后，刷新临时文件的缓冲区，写入临时文件中
	for  {
		if i != n {
			buf,err := reader.ReadString('\n')
			if err == io.EOF{
				writer.Flush()
				break
			}
			writer.WriteString(buf)
			i++
		}else if i == n {
			writer.WriteString(str)
			i++
		}
	}
	// 关闭源文件和临时文件
	f1.Close()
	f2.Close()
	// 重新打开源文件和临时文件
	// 只是这次临时文件是以只读数据的形式打开，源文件是以追加的模式打开
	
	f3,err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Println("file open failed: ", err)
		return
	}
	defer f3.Close()
	f4,err := os.Open("./tmp.txt")
	if err != nil{
		fmt.Println("file open failed: ", err)
		return
	}
	defer f4.Close()
	// 读取临时文件中的数据到缓冲区，为源文件的写入数据创建缓冲区
	reader1 := bufio.NewReader(f4)	
	writer1 := bufio.NewWriter(f3)
	// 逐行读取临时文件中缓冲区的数据，然后写入到源文件的数据缓冲区中直至完全
	for{
		buf1,err := reader1.ReadString('\n')
		if err == io.EOF{
			fmt.Println("insert complet")
			return
		}
		writer1.WriteString(buf1)
		writer1.Flush()
	}
}
// 可以在完成第一步之后使用 os.Rename() 重命名