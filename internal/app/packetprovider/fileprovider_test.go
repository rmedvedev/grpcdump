package packetprovider

import (
	"bytes"
	"testing"
)

func TestFilePacketProvider(t *testing.T) {
	provider, err := NewFileProvider("test/file.pcap")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	chPackets := provider.GetPackets()

	packet := <-chPackets
	if !bytes.Equal(packet.Data(), []byte{1, 2, 3, 4}) {
		t.Error("Not valid data in packet")
	}
}
