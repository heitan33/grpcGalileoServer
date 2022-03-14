package exporter

import (
	pb "proto"
	"fmt"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
)

var address string
var pioResourceInfo *ServerStatItem

func Report (warning bool, serverState, address string) (bool) {
	if warning == true {
	    warning = false
	    pioResourceInfo.Tag = false
	    fmt.Println("FFFFFFFFFF")
	    fmt.Println(pioResourceInfo.Tag)
	    name := serverState
	    conn, err := grpc.Dial(address, grpc.WithInsecure())
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer conn.Close()
	    client := pb.NewHelloClient(conn)
	    request, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	    if err != nil {
	        log.Fatal(err)
	    }
	    fmt.Println(request.Message)
	}
	return warning
}
