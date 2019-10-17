package gogrpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"testing"
	"time"

	ecpb "google.golang.org/grpc/examples/features/proto/echo"
)

func TestNewServer(t *testing.T) {
	os.Setenv(EnvKeyRpcServerAddr, ":9999")
	os.Setenv(EnvKeyRpcServerMode, "release")
	os.Setenv(EnvKeyRpcSSLEnable, "true")
	os.Setenv(EnvKeyRpcSSLServerName, "uufff.com")
	os.Setenv(EnvKeyRpcSSLServerKey, "testdata/cert/server.key")
	os.Setenv(EnvKeyRpcSSLServerCrt, "testdata/cert/server.crt")
	os.Setenv(EnvKeyRpcSSLClientCrt, "testdata/cert/client.crt")
	os.Setenv(EnvKeyRpcRegAddress, "127.0.0.1:9999")
	os.Setenv(EnvKeyRpcRegDomain, "http://uufff.com")
	os.Setenv(EnvKeyRpcOssDomain, "http://uufff.com/attachment")

	// new server
	srv, err := NewServer()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// register
	ecpb.RegisterEchoServer(srv, new(ecServer))

	// run server
	go func() {
		if err := RunServer(srv); err != nil {
			t.Error(err)
			t.FailNow()
		}
	}()

	// sleep
	//time.Sleep(time.Millisecond * 100)
	time.Sleep(time.Second)

	// conn
	conn, err := NewClientWithAddr("127.0.0.1:9999")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// hello world
	res, err := ecpb.NewEchoClient(conn).UnaryEcho(context.Background(), &ecpb.EchoRequest{Message: "hello world."})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("success : %#v \n", res.String())
	fmt.Printf("success : %#v \n", res.String())

	os.Exit(0)
}

// ctrl
type ecServer struct{}

func (s *ecServer) UnaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: req.Message}, nil
}
func (s *ecServer) ServerStreamingEcho(*ecpb.EchoRequest, ecpb.Echo_ServerStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) ClientStreamingEcho(ecpb.Echo_ClientStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) BidirectionalStreamingEcho(ecpb.Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
