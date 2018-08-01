package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//"/"テスト用のハンドラ
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello, World!")
}

//Fizzbuzzのハンドラ
func handleParams(w http.ResponseWriter, r *http.Request){
	num, err := strconv.Atoi(r.URL.Path[10:])
	//条件分岐
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid\n")
	}else if num%15 == 0{
		fmt.Fprintf(w,"Fizzbuzz\n")
	}else if num%5 == 0{
		fmt.Fprintf(w, "buzz\n")
	}else if num%3 == 0{
		fmt.Fprintf(w,"Fizz\n")
	}else {
		fmt.Fprintf(w,"%s\n",strconv.Itoa(num))
	}

	r.ParseForm()
}


func main() {
	http.HandleFunc("/",handler)
	http.HandleFunc("/Fizzbuzz/",handleParams)
	http.ListenAndServe(":8080",nil)
}
