package main

import (
	"fmt"
	"runtime"

	"github.com/containers/gvisor-tap-vsock/pkg/types"
	"github.com/containers/gvisor-tap-vsock/pkg/virtualnetwork"
)

func main() {
	config := types.Configuration{
		Debug:             false,
		CaptureFile:       "",
		MTU:               1500,
		Subnet:            "10.10.10.0/16",
		GatewayIP:         "10.10.10.1",
		GatewayMacAddress: "12:34:56:78:aa:bb",
		DHCPStaticLeases:  map[string]string{},
		DNS:               []types.Zone{},
		Forwards:          map[string]string{},
		NAT:               map[string]string{},
		GatewayVirtualIPs: []string{"10.10.10.254"},
		Protocol:          types.QemuProtocol,
	}

	fmt.Println("Memory usage before creating virtual networks:")
	memStats()

	for i := 0; i < 1000; i++ {
		_, err := virtualnetwork.New(&config)
		if err != nil {
			panic(err)
		}
	}

	runtime.GC()

	fmt.Println("Memory usage after creating virtual networks:")
	memStats()
}

func memStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
