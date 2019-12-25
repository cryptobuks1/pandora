package p2p

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/rs/zerolog"
)

const topic = "pandora"

func pubsubHandler(sub *pubsub.Subscription, logger zerolog.Logger) {
	for {
		msg, err := sub.Next(context.TODO())
		if err != nil {
			logger.Err(err).Msg("get next subscription message error")
			continue
		}

		fmt.Println(msg.Data)
	}
}
