package server

import (
	pb "github.com/baopham/goproto/summarizer"
)

type SummarizerServer struct {
	pb.SummarizerServiceServer
}
