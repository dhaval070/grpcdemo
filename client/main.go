package main

import (
    "fmt"
    "google.golang.org/grpc"
    "grpcdemo/hellopb"
    "context"
    "log"
    "time"
)

func main() {
    conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithInsecure(), grpc.WithBlock())

    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    c := hellopb.NewHelloServiceClient(conn)

    name := "dhaval"

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resp, err := c.Hello(ctx, &hellopb.HelloRequest{ Name: name })

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)

}
