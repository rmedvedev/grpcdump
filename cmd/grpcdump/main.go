package main

import (
	"flag"
	"fmt"

	"github.com/rmedvedev/grpcdump/internal/app/filter"
	"github.com/rmedvedev/grpcdump/internal/app/httpparser"
	"github.com/rmedvedev/grpcdump/internal/app/models"
	"github.com/rmedvedev/grpcdump/internal/app/packetprovider"
	"github.com/rmedvedev/grpcdump/internal/app/protoprovider"
	"github.com/rmedvedev/grpcdump/internal/app/renderers"
	"github.com/rmedvedev/grpcdump/internal/pkg/config"
	"github.com/rmedvedev/grpcdump/internal/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	config.Init()
	cfg := config.GetConfig()

	err := protoprovider.Init(cfg.ProtoPaths, cfg.ProtoFiles)
	if err != nil {
		logrus.Fatal("Proto files init error: ", err)
	}

	err = logger.Init(config.GetConfig().LoggerLevel)
	if err != nil {
		logrus.Fatal("Logger init error: ", err)
	}

	logrus.Infof("Starting sniff ethernet packets at interface %s on port %d", config.GetConfig().Iface, config.GetConfig().Port)

	provider, err := packetprovider.NewEthernetProvider(config.GetConfig().Iface)
	if err != nil {
		logrus.Fatal("Error to create packet provider", err)
	}

	packetFilter := filter.New()
	packetFilter.SetPort(uint32(config.GetConfig().Port))

	err = provider.SetFilter(packetFilter)
	if err != nil {
		logrus.Fatal("Error to create filter", err)
	}

	modelsCh := make(chan models.RenderModel, 1)
	go renderOutput(modelsCh)
	httpParser := httpparser.New(&modelsCh)
	packets := provider.GetPackets()

	for {
		select {
		case packet := <-packets:
			if packet == nil {
				return
			}
			err = httpParser.Parse(packet)
			if err != nil {
				logrus.Warning(err)
			}
		}
	}
}

func renderOutput(models chan models.RenderModel) {
	renderer := renderers.GetApplicationRenderer()
	for {
		select {
		case model := <-models:
			fmt.Println(renderer.Render(model))
		}
	}
}
