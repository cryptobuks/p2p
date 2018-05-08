package main

import (
	"net"

	"github.com/subutai-io/p2p/lib"
)

// Proxy is responsible for keeping track of proxy servers received from bootstrap nodes
// and determining fastest to use by instances

type proxyEntity struct {
	addr    *net.UDPAddr
	latency int
}

type proxy struct {
}

// Request proxies from bootstrap
func (p *proxy) request() error {
	ptp.Log(ptp.Debug, "Requesting proxies")

	packet := &ptp.DHTPacket{
		Type:    ptp.DHTPacketType_Proxy,
		Version: ptp.PacketVersion,
	}

	bootstrap.send(packet)

	return nil
}

// Handle received proxy
func (p *proxy) handleProxyResponse(packet *ptp.DHTPacket) error {

	return nil
}
