package main

import (
	"context"
	"fmt"
	// "log"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	// "github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

const DiscoveryServiceTag = "filego-discovery"

type DiscoveryNotifee struct {
	h host.Host
}

func (n *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo){
	fmt.Printf("Discovered new peer: %s\n", pi.ID.String())

	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("Failed to connect to peer %s: %v\n", pi.ID.String(), err)
	} else {
		fmt.Printf("Successfully connected to peer %s\n", pi.ID.String())
	}
}

func setupDiscovery(h host.Host) {
	mdnsService := mdns.NewMdnsService(h, DiscoveryServiceTag, &DiscoveryNotifee{h : h})
	if err := mdnsService.Start(); err != nil {
		fmt.Printf("Failed to start mDNS service: %v", err)
	}
	fmt.Printf("mDNS discovery service started")
}

func main() {

	host, err := libp2p.New()
	if err != nil {
		fmt.Printf("failed to create libp2p host: %v", err)
	}

	setupDiscovery(host)

	select {}
}