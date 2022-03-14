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
	Email	 string	`yaml:"email"`
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
	var c conf
	conf := c.getConf()
	email := conf.Email
	var rpcPost exporter.ServerStatItem
	json.Unmarshal([]byte(in.Name), &rpcPost)
	post(email, rpcPost, in.Name)
	return &pb.HelloReply{
		Message: "hello " + in.Name,
	}, nil
}


func post(receiveEmail string, contain exporter.ServerStatItem, rpcPost string) {
	client := email.New("yyy3988@qq.com", "qafgnzmwqknucaae", "wbei", "smtp.qq.com", 465, true)
	if contain.Tag == true {
		if err := client.SendEmail([]string{receiveEmail}, "WARNING" + " " + contain.HostName, rpcPost);err !=nil{
			fmt.Println(err)
		}
	} else {
        if err := client.SendEmail([]string{receiveEmail}, "RECOVERY" + " " + contain.HostName, rpcPost);err !=nil{
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

