package collectors

import "context"

// Host on which the dockermon-go is running
var Host HostStaticData

// Static host data that does not change overtime for the given host
type HostStaticData struct {
	HostName      string
	Os            string
	MemoryTotalGB int
	DockerVersion string
}

// Static host data setter
func SetHostStaticData(client doClient) {
	info, _ := client.dc.Info(context.Background())
	Host.HostName = info.Name
	Host.Os = info.OperatingSystem
	Host.MemoryTotalGB = int(info.MemTotal / 1024 / 1024)
	Host.DockerVersion = info.ServerVersion
}

func getHostStaticData() *HostStaticData {
	return &Host
}
