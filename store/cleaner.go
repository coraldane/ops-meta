package store

import (
	"time"
)

func CleanStaleHost() {
	d := time.Duration(3600) * time.Second
	for {
		time.Sleep(d)
		cleanStaleHost()
	}
}

func cleanStaleHost() {
	// one days ago
	before := time.Now().Unix() - 3600*24

	hostNames := HostAgents.Keys()
	for _, hostname := range hostNames {
		has, exists := HostAgents.Get(hostname)
		if !exists {
			continue
		}
		if has.Timestamp <= before {
			HostAgents.Delete(hostname)
		}
	}
}
