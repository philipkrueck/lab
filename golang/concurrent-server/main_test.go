package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type sku struct {
	item, price string
}

var items = []sku{
	{"socks", "6"},
	{"sandals", "27"},
	{"pants", "30"},
	{"shorts", "20"},
	{"shirt", "32"},
	{"jacket", "79"},
	{"cap", "22"},
}

func sendRequest(cmd, params string) {
	resp, err := http.Get("http://localhost:8081/" + cmd + params)
	if err != nil {
		log.Printf("params = %s; err: %s\n", params, err)
		return
	}

	defer resp.Body.Close()
	log.Printf("got %s = %d (err: %s)\n", params, resp.StatusCode, err)
}

func runCreates() {
	for {
		for _, sku := range items {
			params := fmt.Sprintf("?item=%s&price=%s", sku.item, sku.price)
			sendRequest("create", params)
		}
	}
}

func runUpdates() {
	for {
		for _, sku := range items {
			params := fmt.Sprintf("?item=%s&price=%s", sku.item, sku.price)
			sendRequest("update", params)
		}
	}
}

func runDeletes() {
	for {
		for _, sku := range items {
			params := fmt.Sprintf("?item=%s", sku.item)
			sendRequest("delete", params)
		}
	}
}

func TestServer(t *testing.T) {
	go runCreates()
	go runUpdates()
	go runDeletes()

	time.Sleep(5 * time.Second)
}
