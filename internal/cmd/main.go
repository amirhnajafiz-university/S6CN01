package cmd

import (
	"github.com/amirhnajafiz/packet-monitoring/internal/agent"
	"github.com/amirhnajafiz/packet-monitoring/internal/client"
	"github.com/amirhnajafiz/packet-monitoring/internal/config"
)

func Execute() {
	// load configs
	cfg := config.Load()

	// starting prometheus client
	go client.Client{}.Start(cfg.Client)

	// starting our agents
	for i := 0; i < cfg.NumberOfAgents; i++ {
		go agent.Agent{}.Start(cfg.Agent)
	}

	// busy waiting
	select {}
}
