package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"reflect"
)

// ini 配置解析
type MysqlConfig struct{
	Address string `ini:"address"`
	Port int `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
// 将字符串转换为 int 
func StringToInt(str string) int {
	var b = []byte(str)
	var value, sum int
	for i,v := range b{
		value = int(v-'0')
		for j := 0; j < len(b)-i-1; j++{
			value *= 10
		}
		sum += value
	}
	return sum
}

func LoadIni(fileName string, data interface{})(err error) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	_ = v
	// 判断 data 参数是否是结构体指针
	if t.Kind() != reflect.Ptr{
		err = fmt.Errorf("data should be a pointer")
		return
	}
	if t.Elem().Kind() != reflect.Struct{
		err = errors.New("data param should be struct")
	}
	// 读取配置文件内容
	b,err1 := ioutil.ReadFile(fileName)
	if err1 != nil{
		fmt.Println(err1)
		return
	}
	// 对配置文件内容进行格式化
	lineSlice0 := strings.Split(string(b), "\r\n")
	fmt.Println(lineSlice0)
	for idx,line0 := range lineSlice0{
		// 去除字符串两端的空格
		line1 := strings.TrimSpace(line0)
		// 去除空行
		if len(line1) == 0 {
			continue
		}
		// 去除注释 ";" "#"
		if strings.HasPrefix(line1,";") || strings.HasPrefix(line1,"#"){
			continue
		}
		// 如果以 "["开头，则表示是 section
		if strings.HasPrefix(line1,"[") {
			if !strings.HasSuffix(line1,"]"){
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 提取 section name
			sectionName := line1[1:len(line1)-1]
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
		} else {
			// 判断是否有 "="，如果有 "="，则将字符串以 "=" 进行切割，
			if strings.Index(line1, "=") == -1 || strings.HasPrefix(line1, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// "=" 左边的为 key, "=" 右边的为 value
			index := strings.Index(line1, "=")
			key := strings.TrimSpace(line1[:index])
			value := strings.TrimSpace(line1[index+1:])
		
			for i := 0; i < t.Elem().NumField(); i++ {
				fName := t.Elem().Field(i).Name // fName 为结构体的字段名
				fPtr := v.Elem().FieldByName(fName).Addr() // fPtr 为结构体字段名所对应的指针
				sKey := t.Elem().Field(i).Tag.Get("ini") // sKey 为指定格式所对应的字段的字符串
				sType := t.Elem().Field(i).Type  // sType 为结构体指定字段的类型
				if sKey == key {
					if sType.Kind() == reflect.Int {  // 如果指定的字段是 Int 类型则需将字符串先转换成 Int 
						sValue := StringToInt(value)
						fPtr.Elem().SetInt(int64(sValue))
					} else if sType.Kind() == reflect.String {
						fPtr.Elem().SetString(value)
					}
				}
			}
		}
	}
	return 
}

func main(){
	var mc = MysqlConfig{}
	LoadIni("./mysqlconfig.ini", &mc)
	fmt.Println(mc)
}