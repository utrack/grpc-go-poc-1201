# grpc-go #1201 POC
This POC demonstrates erroneous behaviour of gRPC client when using
1.7-unaware net/context library.

While current net/context proxies calls to vanilla `context`, pre-1.7
version has its own context.Context type and errors (like `netctx.Canceled`
et al) that do not equal `"context".Canceled`.

gRPC client panics if it sees context.Canceled (because it expects
`net/context.Canceled`).
Client recovers from panic and runs as if nothing happened.
Some time later, client becomes "stuck" - it does not send any requests
to the server, even with perfectly good context and request.

Requests on client side are stuck indefinitely.

There's some other corner case that causes server to leak routines
in similar situation, but I didn't manage to catch it.

## Running
`go run ./cmd/server/main.go`

In other console:
`go run ./cmd/client/main.go -cancel=true -panic=true`
