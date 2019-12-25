package p2p

import (
	"bufio"
	"context"
	"crypto/rand"
	"os"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/rs/zerolog"
)

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

func chatInputLoop(topic *pubsub.Topic, logger zerolog.Logger, doneC chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		send(msg, topic, logger)
	}

	doneC <- struct{}{}
}
