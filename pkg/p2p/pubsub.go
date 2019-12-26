package p2p

import (
	"bufio"
	"context"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/rs/zerolog"
)

const topic = "pandora"

func pubSubHandler(sub *pubsub.Subscription, logger zerolog.Logger) {
	for {
		msg, err := sub.Next(context.TODO())
		if err != nil {
			logger.Err(err).Msg("get next subscription message error")
			continue
		}

		fmt.Println(string(msg.Data))
	}
}

func send(msg string, topic *pubsub.Topic, logger zerolog.Logger) {
	var err error

	defer func() {
		if err != nil {
			logger.Err(err).Msg("send message error")
		}
	}()

	msgID := make([]byte, 10)
	_, err = rand.Read(msgID)

	err = topic.Publish(context.TODO(), []byte(msg))
}

func inputLoop(topic *pubsub.Topic, logger zerolog.Logger, doneC chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		send(msg, topic, logger)
	}
	doneC <- struct{}{}
}
