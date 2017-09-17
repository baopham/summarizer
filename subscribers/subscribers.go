package subscribers

import (
	pb "github.com/baopham/gominderproto"
	"github.com/lileio/lile/pubsub"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// See https://godoc.org/github.com/lileio/lile/pubsub#Handler for an example
// of what an subscriber handler should look like

type SummarizerServiceSubscriber struct{}

func (s *SummarizerServiceSubscriber) Setup(c *pubsub.Client) {
	// https://godoc.org/github.com/lileio/lile/pubsub#Client.On
	c.On("summarizer.start", s.Summarize, 30*time.Second, true)
}

func (s *SummarizerServiceSubscriber) Summarize(ctx context.Context, r *pb.SummarizerRequest) {
	log.Printf("got event %s", r)
	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewSummarizerServiceClient(conn)

	summary, err := serviceClient.Summarize(ctx, r)
	log.Printf("%v %v", summary, err)
}
