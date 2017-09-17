package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/baopham/gominderproto"
	"github.com/lileio/lile"
)

var s = SummarizerServer{}
var cli pb.SummarizerClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		pb.RegisterSummarizerServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = pb.NewSummarizerClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
