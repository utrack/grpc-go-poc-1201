package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	pb "github.com/utrack/grpcgo-poc-1201/pb"
	"google.golang.org/grpc"

	"context"

	netctx "golang.org/x/net/context"
)

// This POC demonstrates erroneous behaviour of gRPC client when using
// 1.7-unaware net/context library.
//
// While current net/context proxies calls to vanilla "context", pre-1.7
// version has its own context.Context type and errors (like netctx.Canceled
// et al) that do not equal "context".Canceled.
//
// gRPC client panics if it sees context.Canceled (because it expects
// net/context.Canceled).
// Client recovers from panic and runs as if nothing happened.
// Some time later, client becomes "stuck" - it does not send any requests
// to the server, even with perfectly good context and request.
//
// Requests on client side are stuck indefinitely.
//
// There's some other corner case that causes server to leak routines
// in similar situation, but I didn't manage to catch it.

func main() {
	logrus.Info("Client starting")

	conn, err := grpc.Dial("127.0.0.1:1201", grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Connected")
	cli := pb.NewSummatorClient(conn)

	// At first, run requests with bad net/context and cancel ctx
	// when awaiting rsp from server
	doPanic, doCancel := true, true
	panickedCalls := 0
	for {
		var ctx context.Context
		var cancelFunc func()
		if doPanic {
			ctx, cancelFunc = context.WithCancel(context.Background())
			panickedCalls++
			// Stop using net/context after 500 requests just to be sure
			// After 500 reqs requests shouldn't panic and start
			// working normally
			if panickedCalls > 500 {
				doPanic = false
				doCancel = false
			}
		} else {
			ctx, cancelFunc = netctx.WithCancel(netctx.Background())
		}

		// Cancel requests while we sent them but not rcvd
		if doCancel {
			go func() {
				<-time.After(time.Millisecond * 500)
				cancelFunc()
			}()
		}
		go makeCall(ctx, cli)
		<-time.After(time.Millisecond * 15)
	}
}

func makeCall(ctx context.Context, cli pb.SummatorClient) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Warn("recovered")
		}
	}()
	if _, err := cli.Sum(ctx, &pb.SumRequest{A: 1, B: 2}); err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info("ok")
}
