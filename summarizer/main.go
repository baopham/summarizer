package main

import (
	pb "github.com/baopham/goproto/summarizer"
	"github.com/baopham/summarizer/server"
	"github.com/baopham/summarizer/summarizer/cmd"
	"github.com/lileio/lile"
	"google.golang.org/grpc"
)

func main() {
	s := &server.SummarizerServer{}

	lile.Port(":8001")
	lile.Name("summarizer")
	lile.Server(func(g *grpc.Server) {
		pb.RegisterSummarizerServiceServer(g, s)
	})

	cmd.Execute()
}
