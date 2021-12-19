package main

import (
    "fmt"
    "net"
    "google.golang.org/grpc"
    "grpcdemo/hellopb"
    "context"
    "log"
)

type server struct {
    hellopb.UnimplementedHelloServiceServer

}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {

    name := req.Name
    response := &hellopb.HelloResponse {
        Greeting: "hello " + name,
    }

    return response, nil
}

func main() {
    address := "0.0.0.0:9000"
    
    sock, err := net.Listen("tcp", address) 

    if err != nil {
        log.Fatal(err)
    }

    s := grpc.NewServer()

    hellopb.RegisterHelloServiceServer(s, &server{})

    fmt.Println("listenning...")
    s.Serve(sock)
}
