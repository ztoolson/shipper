package main

import (
	"fmt"
	"log"
	"net"

	// import the generated protobuf code
	pb "github.com/ztoolson/shipper/consignment-service/proto/consignment"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Hello World")
}
