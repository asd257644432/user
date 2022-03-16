package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"user/kitex_gen/user/user"
)

func main() {
	ipAddr, err := net.ResolveIPAddr("tcp", "127.0.0.1:8000")
	if err != nil {
		panic("failed to lesson port")
	}
	svr := user.NewServer(new(UserImpl), server.WithServiceAddr(ipAddr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
