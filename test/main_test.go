package test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	failed := false

	node := newNode()

	if err := node.start(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	_, err := http.Get("http://127.0.0.1:9999/live")
	if err != nil {
		failed = true
	}

	time.Sleep(time.Second * 5)

	if err := node.stop(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	if failed {
		os.Exit(1)
	}
	os.Exit(0)
}
