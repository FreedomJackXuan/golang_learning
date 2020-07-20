package main

import (
	"fmt"
	"golang_learning/src/zinx/znet"
	"net"
	"testing"
	"time"
)

func ClientTest(){
	fmt.Println("Client Test start")
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("connection error", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Println("write error: ",err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}
		fmt.Println("server call back: ", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}

func TestClient(t *testing.T){
	//fmt.Println("xxx")
	s := znet.NewServer("[zinx V0.1]")
	go ClientTest()
	s.Serve()
}
