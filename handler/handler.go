package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
	"os"
)

func UploadHandler(w http.ResponseWriter,r *http.Request) {
	if r.Method=="GET"{
		//返回上传界面
		data,err:=ioutil.ReadFile("./static/view/index.html")
		if err!=nil{
			io.WriteString(w,"Inter Server Error")
			return
		}
		io.WriteString(w,string(data))
	}else if r.Method=="POST"{
		//接受文件流到本地
		file, header, err := r.FormFile("file")
		if err!=nil{
			fmt.Printf("Fail to create file,err:%s",err.Error())
			return
		}
		//关闭文件流
		defer file.Close()

		newFile,err := os.Create("./tmp/" + header.Filename)
		if err!=nil{
			fmt.Printf("Fail to create file,err:%s",err.Error())
			return
		}
		defer newFile.Close()

		_,err = io.Copy(newFile, file)
		if err!=nil{
			fmt.Printf("Fail to save data to file,err:%s",err.Error())
			return
		}

		http.Redirect(w,r,"/file/upload/success",http.StatusFound)
	}
}

func UploadFileSucHandler(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"upload successfully")
}