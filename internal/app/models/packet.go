package models

import (
	"fmt"

	"github.com/google/gopacket"
)

//Packet represents entity net request for view
type Packet struct {
	NetSrc       string
	NetDst       string
	TransportSrc string
	TransportDst string
}

//NewPacket creates new request view from gopacket.Flow
func NewPacket(net, transport gopacket.Flow) *Packet {

	return &Packet{
		NetSrc:       net.Src().String(),
		NetDst:       net.Dst().String(),
		TransportSrc: transport.Src().String(),
		TransportDst: transport.Dst().String(),
	}
}

//GetConnectionKey returns connection string
func (packet *Packet) GetConnectionKey() string {
	return fmt.Sprintf("%s:%s->%s:%s", packet.NetSrc, packet.TransportSrc, packet.NetDst, packet.TransportDst)
}

//GetRevConnectionKey returns connection string
func (packet *Packet) GetRevConnectionKey() string {
	return fmt.Sprintf("%s:%s->%s:%s", packet.NetDst, packet.TransportDst, packet.NetSrc, packet.TransportSrc)
}
