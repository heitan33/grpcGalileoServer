package main

import (
    "golang.org/x/net/context"
    pb "proto"
    "net"
    "log"
    "google.golang.org/grpc"
    "fmt"
	"grpcServer/email"
	"strings"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"encoding/json"
	"grpcServer/exporter"
)


type conf struct {
  Host     string `yaml:"host"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("galileo.yaml")
	if err != nil {
		log.Println("yamlFile.Get err", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println("Unmarshal: ", err.Error())
	}
	return c
}

type Server struct {}
type ServerStatItem exporter.ServerStatItem
type contain exporter.ServerStatItem


func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	var rpcPost exporter.ServerStatItem
	json.Unmarshal([]byte(in.Name), &rpcPost)
	post(rpcPost, in.Name)
	return &pb.HelloReply{
		Message: "hello " + in.Name,
	}, nil
}


func post(contain exporter.ServerStatItem, rpcPost string) {
	client := email.New("yyy3988@qq.com", "qafgnzmwqknucaae", "wbei", "smtp.qq.com", 465, true)
	if contain.Tag == true {
		if err := client.SendEmail([]string{"yyy3988@qq.com"}, "WARNING", rpcPost);err !=nil{
			fmt.Println(err)
		}
	} else {
        if err := client.SendEmail([]string{"yyy3988@qq.com"}, "RECOVERY", rpcPost);err !=nil{
            fmt.Println(err)
        }
	}
}

func main() {
	var c conf
	conf := c.getConf()
	port := strings.Split(conf.Host, ":")[1]
	port = ":" + port

	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc server listening at: 50051 port")
	server := grpc.NewServer()
	pb.RegisterHelloServer(server, &Server{})
	server.Serve(conn)
}

