package packetprovider

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"github.com/rmedvedev/grpcdump/internal/app/filter"
)

//EthernetProvider ...
type EthernetProvider struct {
	handler *pcapgo.EthernetHandle
}

//NewEthernetProvider create new EthernetProvider
func NewEthernetProvider(iface string) (PacketProvider, error) {
	handler, err := pcapgo.NewEthernetHandle(iface)
	if err != nil {
		return nil, err
	}

	return &EthernetProvider{handler}, nil
}

//SetFilter sets bpf filter
func (provider *EthernetProvider) SetFilter(packetFilter *filter.PacketFilter) (err error) {
	err = provider.handler.SetBPF(packetFilter.GetBpfFilter())
	return
}

//GetPackets return channel for get packets
func (provider *EthernetProvider) GetPackets() chan gopacket.Packet {
	packetSource := gopacket.NewPacketSource(provider.handler, layers.LayerTypeEthernet)
	return packetSource.Packets()
}
