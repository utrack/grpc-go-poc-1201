package main

import (
	"fmt"
	"time"

	"net"

	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/soheilhy/cmux"
	pb "github.com/utrack/grpcgo-poc-1201/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type SumImpl struct{}

var _ pb.SummatorServer = &SumImpl{}

func (s *SumImpl) Sum(ctx context.Context, r *pb.SumRequest) (*pb.SumResponse, error) {
	//log.Println("Processed")
	<-time.After(time.Second)
	fmt.Print(".")
	return &pb.SumResponse{Result: r.GetA() + r.GetB()}, nil
}

func main() {
	logrus.Info("Starting")
	go goroutLoop()
	srv := grpc.NewServer()
	lis, err := net.Listen("tcp", "127.0.0.1:1201")
	if err != nil {
		logrus.Fatal(err)
	}
	cmlis := cmux.New(lis)
	glis := cmlis.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	go cmlis.Serve()
	pb.RegisterSummatorServer(srv, &SumImpl{})
	err = srv.Serve(glis)
	if err != nil {
		logrus.Fatal(err)
	}
}

func goroutLoop() {
	for {
		logrus.Printf("goroutines count: %v", runtime.NumGoroutine())
		<-time.After(time.Second * 2)
	}
}
