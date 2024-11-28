package general

import "github.com/guneyin/locator/util"

type (
	ServiceStatus string
	Env           string
)

const (
	ServiceStatusOnline      ServiceStatus = "online"
	ServiceStatusMaintenance ServiceStatus = "maintenance"

	EnvProduction Env = "production"
	EnvStaging    Env = "staging"
)

type Status struct {
	Status  ServiceStatus
	Version *util.VersionInfo
	Env     Env
	Uptime  string
}
