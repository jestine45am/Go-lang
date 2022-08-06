package main

import (
	"fmt"
	"io/ioutil"
	// "io/ioutil"
	// "net/http"
	"net/http"
	"net/url"
)

func main() {
/* 	resp,err := http.Get("http://localhost:8080/hello")
	if err != nil {
		fmt.Println("Get url failed:", err)
		return
	}
	fmt.Println(resp.Status, resp.Header)
	defer resp.Body.Close()
	respBody,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}
	fmt.Println(string(respBody))
 */
	
	apiurl := "http://localhost:8080/index.html"
	url := url.Values{}
	url.Add("name", "zhangsan")
	url.Set("age", "18")
	newurl := apiurl + "?" + url.Encode()
	resp, err := http.Get(newurl)
	if err != nil {
		fmt.Println("Get url failed:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

}