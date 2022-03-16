package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"user/kitex_gen/user/user"
)

func main() {
	ipAddr, _ := net.ResolveTCPAddr("tcp", ":8000")
	svr := user.NewServer(new(UserImpl), server.WithServiceAddr(ipAddr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
