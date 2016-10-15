package collectors

import "context"

var Host HostStaticData

type HostStaticData struct {
	HostName          string
	Os                string
	MemoryTotalGB     int
	DockerVersion     string
}

func SetHostStaticData(client doClient)  {
	info, _ := client.dc.Info(context.Background())
	Host.HostName = info.Name
	Host.Os = info.OperatingSystem
	Host.MemoryTotalGB = int (info.MemTotal / 1024 / 1024)
	Host.DockerVersion = info.ServerVersion
}

func getHostStaticData() *HostStaticData {
	return &Host
}
