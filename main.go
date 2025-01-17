package main

import (
	"fmt"
	"net/http"
	"oss-netdisk/handler"
)

func main() {
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/success",handler.UploadFileSucHandler)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Printf("Fail to start server,err:%s",err.Error())
	}
}