/*
Created on 2018/5/30 17:37

author: ChenJinLong

Content: 
*/
package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"io"
	"html/template"
)


func sayhelloName(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()    //解析url传递的参数，对于Post则解析响应包的主体
	//没有这个方法，下面无法获取表单的数据
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	reStr := "hello"
	for k,v := range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
		reStr = reStr + strings.Join(v,",")
	}
	io.WriteString(w,reStr)

}



func login(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		r,_ := template.ParseFiles("login.html")
		log.Println(r.Execute(w,nil))
	}else {
		fmt.Println("username",r.Form["username"])
		fmt.Println("password",r.Form["password"])

	}

}





func main() {

	http.HandleFunc("/login",login)
	http.HandleFunc("/",sayhelloName)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)

	}


}
