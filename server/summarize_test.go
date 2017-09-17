package server

import (
	"testing"

	pb "github.com/baopham/gominderproto"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestSummarize(t *testing.T) {
	ctx := context.Background()
	req := &pb.SummarizerRequest{}

	res, err := cli.Summarize(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
