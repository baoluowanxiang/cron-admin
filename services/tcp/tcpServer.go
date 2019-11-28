package tcp

import (
	"cron-admin/option"
	"fmt"
	"log"
	"net"
)

type TcpServer struct {

}

func (t *TcpServer) Start(opts option.Options) error {
	listener, err := net.Listen("tcp", "0.0.0.0:9600")
	if err != nil {
		return nil
	}
	for {
		conn, err := listener.Accept() //开启监听
		if err != nil {
			fmt.Println("Accept is err!: ", err)
			continue
		}
		t.dealWithConnection(conn)
	}
}

func (t *TcpServer) dealWithConnection(connection net.Conn) {
	log.Print(connection.RemoteAddr())
}