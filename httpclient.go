package main

import (
	"net/http"
	"fmt"
	"net/http/httputil"
	"bytes"
	"strconv"
)

//1-MaxNumまでFizzbuzzする
var MaxNum int = 60

//13の倍数で文字列が含まれたURLを作成する関数
func jadgeNum(num int)string{

	if num%13 == 0{
		return "http://localhost:8080/Fizzbuzz/!!Enjoing!!SecurityCamp2018!!"

	}else { //通常は数値をパラメータとして渡す関数
		buf := bytes.NewBufferString("http://localhost:8080/Fizzbuzz/")
		buf.WriteString(strconv.Itoa(num))
		return  buf.String()
	}
}

func main(){

	for num:=0;num<=MaxNum;num++ {

		url := jadgeNum(num)

		//リクエストを投げる
		request, _ := http.NewRequest("Get",url, nil)
		client := new(http.Client)
		response, _ := client.Do(request)


		//レスポンスヘッダを取得
		dump, _ := httputil.DumpResponse(response, true)

		//出力
		fmt.Println("===========================================")
		fmt.Printf("\n curl -v %s \n \n %s \n", url, dump)
		fmt.Println("===========================================")
	}
}
