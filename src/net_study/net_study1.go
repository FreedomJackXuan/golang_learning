package net_study

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func NetStudy1(){
	url := "json.cn"
	prefix := "http://"
	if !strings.HasPrefix(url, prefix) {
		url = prefix + url
	}
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch1: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(response.Status)
	fmt.Println(response.Header)
	buf, err := io.Copy(os.Stdout, response.Body)
	//buf, err := ioutil.ReadAll(response.Body)
	response.Body.Close() // 需要显式释放系统层资源，Go虽然有垃圾回收，将不使用的内存释放，但是不包括操作系统层面的资源

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch2: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s",buf)

}
