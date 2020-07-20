package znet

import (
	"fmt"
	"golang_learning/src/zinx/ziface"
	"net"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	handleAPI ziface.HandFunc
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandFunc) * Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		isClosed: false,
		handleAPI: callback_api,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

func (self *Connection) StartReader() {
	fmt.Println("Reader groutine is running")
	defer fmt.Println(self.RemoteAddr().String(), " conn reader exit!")
	defer self.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := self.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err: ", err)
			self.ExitBuffChan <- true
			continue
		}
		if err := self.handleAPI(self.Conn, buf, cnt); err != nil {
			fmt.Println("connID ", self.ConnID, " handler is err")
			self.ExitBuffChan <- true
			return
		}
	}
}

func (self *Connection) Start() {
	go self.StartReader()
	for {
		select {
		case <- self.ExitBuffChan:
			return
		}
	}
}

func (self *Connection) Stop() {
	if self.isClosed == true {
		return
	}
	self.isClosed = true

	_ = self.Conn.Close()
	self.ExitBuffChan <- true

	close(self.ExitBuffChan)
}

func (self *Connection) GetTCPConnection() *net.TCPConn {
	return self.Conn
}

func (self *Connection) RemoteAddr() net.Addr {
	return self.Conn.RemoteAddr()
}