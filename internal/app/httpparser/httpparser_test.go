package httpparser

import (
	"testing"

	"github.com/rmedvedev/grpcdump/internal/app/models"
	"github.com/rmedvedev/grpcdump/internal/app/packetprovider"
	"github.com/rmedvedev/grpcdump/internal/app/protoprovider"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGrpcHttpParser(t *testing.T) {
	provider, err := packetprovider.NewFileProvider("test/grpc.pcap")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	modelsCh := make(chan models.RenderModel, 100)

	httpParser := New(&modelsCh)
	packets := provider.GetPackets()

	for packet := range packets {
		if packet == nil {
			return
		}

		err = httpParser.Parse(packet)
		if err != nil {
			logrus.Warning(err)
		}
	}

	if len(modelsCh) != 10 {
		t.Errorf("Incorrect number of packet: %d", len(modelsCh))
		t.FailNow()
	}

	grpcRequest := <-modelsCh
	grpcResponse := <-modelsCh

	assert.Equal(t, "/helloworld.Greeter/SayHello", grpcRequest.GetPath())
	assert.Equal(t, "55853", grpcRequest.GetSrcPort())
	assert.Equal(t, "50051", grpcRequest.GetDstPort())
	assert.Equal(t, "50051", grpcResponse.GetSrcPort())
	assert.Equal(t, "55853", grpcResponse.GetDstPort())
}

func TestGrpcParser(t *testing.T) {
	provider, err := packetprovider.NewFileProvider("test/grpc.pcap")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	protoprovider.Init("./test", []string{"helloworld.proto"})

	modelsCh := make(chan models.RenderModel, 100)

	httpParser := New(&modelsCh)
	packets := provider.GetPackets()

	for packet := range packets {
		if packet == nil {
			return
		}

		err = httpParser.Parse(packet)
		if err != nil {
			logrus.Warning(err)
		}
	}

	if len(modelsCh) != 10 {
		t.Errorf("Incorrect number of packet: %d", len(modelsCh))
		t.FailNow()
	}

	grpcRequest := <-modelsCh
	grpcResponse := <-modelsCh

	assert.Equal(t, "name:\"world\"", grpcRequest.GetBody())
	assert.Equal(t, "message:\"Hello world\"", grpcResponse.GetBody())
}
