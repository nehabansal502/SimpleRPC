package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type TimeServer int64

func (t *TimeServer) GetServerTime(args *Args, reply *int64) error {
	*reply = int64(time.Now().Unix())
	return nil
}

func main() {
	timeServer := new(TimeServer)
	err := rpc.Register(timeServer)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(l, nil)
}
