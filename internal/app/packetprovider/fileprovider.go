package packetprovider

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"github.com/rmedvedev/grpcdump/internal/app/filter"
)

//FileProvider ...
type FileProvider struct {
	reader *pcapgo.Reader
}

//NewFileProvider ...
func NewFileProvider(filename string) (PacketProvider, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %s", err.Error())
	}

	pcapgoReader, err := pcapgo.NewReader(reader)
	if err != nil {
		return nil, fmt.Errorf("Failed to create pcapgo reader: %s", err.Error())
	}

	return &FileProvider{pcapgoReader}, nil
}

//SetFilter sets bpf filter
func (provider *FileProvider) SetFilter(packetFilter *filter.PacketFilter) (err error) {
	//we cant attach bpf filter
	return
}

//GetPackets return channel for get packets
func (provider *FileProvider) GetPackets() chan gopacket.Packet {
	packetSource := gopacket.NewPacketSource(provider.reader, layers.LayerTypeLoopback)
	return packetSource.Packets()
}
