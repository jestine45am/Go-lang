package main

// net/http/server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MyHandler(filename string){
	http.HandleFunc(string("/" + filename), func(w http.ResponseWriter, r *http.Request) {
		b,err := ioutil.ReadFile(string("./" + filename))
		if err != nil {
			w.Write([]byte("Error"))
		}
		w.Write(b)
		fmt.Println(r)
		fmt.Println(r.Host, r.URL, r.URL.RawQuery)
	})
}

func main(){
	var str string 
	// open all files in the current directory
	// fs := []fs.FileInfo{}
	fs, err :=ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Open dir error:", err)
	}
	for _, f := range fs {
		str = f.Name()
		switch str {
			case "index.html": 
				MyHandler(str)
			case "hello": 
				MyHandler(str)
		}
	}
	// listen on port 8080
	http.ListenAndServe("0.0.0.0:8080", nil) 
}