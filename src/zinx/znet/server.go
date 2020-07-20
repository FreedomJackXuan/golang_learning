package znet

import (
	"errors"
	"fmt"
	"golang_learning/src/zinx/ziface"
	"net"
	"time"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port int
}

func (self *Server) Start() {
	fmt.Println("server start", self.Name, self.IP, self.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(self.IPVersion, fmt.Sprintf("%s:%d", self.IP, self.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		listener, err := net.ListenTCP(self.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp error", err)
			return
		}

		fmt.Println("start zinx server")

		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid ++

			go dealConn.Start()
		}
	}()
}

func (self *Server) Stop() {
	fmt.Println("zinx stop")
}

func (self *Server) Serve(){
	self.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Connection Handler] CallBackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func NewServer (name string) ziface.IServer {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 7777,
	}
	return s
}

