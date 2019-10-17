package gogrpc

import (
	"context"
	"google.golang.org/grpc/peer"
	"strings"
)

// ClientIP ip
func ClientIP(ctx context.Context) string {
	if pr, ok := Peer(ctx); ok {
		return pr.Addr.String()[0:strings.LastIndex(pr.Addr.String(), ":")]
	}
	return ""
}

// ClientAddress ip:host
func ClientAddress(ctx context.Context) string {
	if pr, ok := Peer(ctx); ok {
		return pr.Addr.String()
	}
	return ""
}

// Peer peer
func Peer(ctx context.Context) (*peer.Peer, bool) {
	return peer.FromContext(ctx)
}

// Value value
func Value(ctx context.Context, key interface{}) interface{} {
	return ctx.Value(key)
}
