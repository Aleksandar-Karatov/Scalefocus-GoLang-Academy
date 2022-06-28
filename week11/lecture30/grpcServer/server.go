package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"week11Lecture30Task/protoserv"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	protoserv.UnimplementedDataGetterServer
}

func (s *server) ShowData(ctx context.Context, in *protoserv.DataRequest) (*protoserv.DataReply, error) {
	log.Println("Recieved data: title: ", in.GetTitle(), " score: ", in.GetScore())
	return &protoserv.DataReply{Title: "Recieved: " + in.GetTitle(), Score: in.Score}, nil
}
func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Printf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	protoserv.RegisterDataGetterServer(s, &server{})
	log.Printf("Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve %v", err)
	}

}
