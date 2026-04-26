package main

import (
	"flag"
	"stress-test/adapter"
)

func main() {
	url := flag.String("url", "https://example.com", "URL para acessar")
	requests := flag.Int("requests", 1, "Número de requisições a serem realizadas")
	concurrency := flag.Int("concurrency", 1, "Quantidade de chamadas simultaneas")
	flag.Parse()

	adapter.PrepareRequest(*url, *requests, *concurrency)
}
