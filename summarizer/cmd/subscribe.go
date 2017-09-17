package cmd

import (
	"github.com/baopham/summarizer/subscribers"
	"github.com/lileio/lile"
	"github.com/spf13/cobra"
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to and process queue messages",
	Run: func(cmd *cobra.Command, args []string) {
		lile.Subscribe(&subscribers.SummarizerServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(subscribeCmd)
}
