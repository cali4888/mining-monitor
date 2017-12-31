package main

import (
	"github.com/cali4888/mining-monitor/monitors"
	"github.com/cali4888/mining-monitor/strategies"
	"log"
)

func main() {
	log.Println("Hello from raspberry")
	monitor := &monitors.Reachability{}

	rs := strategies.RouterRecovery{}
	rs.Run(monitor)

}
