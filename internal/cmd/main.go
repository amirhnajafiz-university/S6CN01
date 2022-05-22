package cmd

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/packet-monitoring/internal/agent"
	"github.com/amirhnajafiz/packet-monitoring/internal/client"
	"github.com/amirhnajafiz/packet-monitoring/internal/config"
)

func Execute() {
	// load configs
	cfg := config.Load()

	// starting prometheus client
	fmt.Println("Client started ...")
	go client.Client{}.Start(cfg.Client)

	// starting our agents
	for i := 0; i < cfg.NumberOfAgents; i++ {
		fmt.Printf("[%d/%d]Agent started ...\n", i+1, cfg.NumberOfAgents)

		a := agent.Agent{
			ID:       int(time.Now().UnixMilli()),
			BusyTime: 2,
			Util:     agent.Util{},
		}
		a.Start(cfg.Agent)
	}

	// busy waiting
	select {}
}
