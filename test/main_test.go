package test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"pandora/test/node"
)

func TestMain(m *testing.M) {
	failed := false

	nod := node.NewNode()

	if err := nod.Start(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	_, err := http.Get("http://127.0.0.1:9999/live")
	if err != nil {
		failed = true
	}

	time.Sleep(time.Second * 5)

	if err := nod.Stop(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	if failed {
		os.Exit(1)
	}
	os.Exit(0)
}
