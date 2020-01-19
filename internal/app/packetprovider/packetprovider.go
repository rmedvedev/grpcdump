package packetprovider

import (
	"github.com/google/gopacket"
	"github.com/rmedvedev/grpcdump/internal/app/filter"
)

//PacketProvider ...
type PacketProvider interface {
	SetFilter(packetFilter *filter.PacketFilter) (err error)
	GetPackets() chan gopacket.Packet
}
