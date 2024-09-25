package main

import (
	"github.com/kaliban2056/Decentralized-poker/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)

	server.Start()

	//rand.NewSource(time.Now().UnixNano())

	//d := deck.New()

	//fmt.Println(d)
}
