package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main(){

	http.HandleFunc("/count", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
	//testDiscard()
	/*var start = time.Now()
	var ch = make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}
	for range os.Args[1:]{
		fmt.Println(<- ch)
	}
	fmt.Println(time.Since(start).Seconds())*/
}

func fetch(url string, ch chan <- string){
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprint("reading " + url)
		return
	}
	ch <- fmt.Sprintf("%7d, %s", nbytes, url)
}

func testDiscard(){
	a := strings.NewReader("hello")
	nbytes, err := io.Copy(io.Discard, a)
	if err != nil{
		panic(err)
	}
	fmt.Println(a)
	fmt.Println(nbytes)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// 解析的所有参数
	if err := r.ParseForm(); err != nil{
		log.Print(err)
	}

	for k, v := range r.Form{
		fmt.Println("k:", k, "\tv:", v)
	}
}