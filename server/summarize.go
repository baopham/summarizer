package server

import (
	pb "github.com/baopham/gominderproto"
	context "golang.org/x/net/context"
)

func (s SummarizerServer) Summarize(ctx context.Context, r *pb.SummarizerRequest) (*pb.Summary, error) {
	return &pb.Summary{Tags: []string{"hello", "world"}}, nil
}
