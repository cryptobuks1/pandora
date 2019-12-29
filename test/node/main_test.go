package main

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	node := NewNode()

	if err := node.Start(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	_, _ = http.Get("http://127.0.0.1:4001/live")

	time.Sleep(time.Second * 5)

	if err := node.Stop(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)
}
