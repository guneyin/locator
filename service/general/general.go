package general

import (
	"github.com/guneyin/locator/util"
	"time"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Status() Status {
	uptime := time.Now().Sub(util.GetLastRun())

	return Status{
		Status:  ServiceStatusOnline,
		Version: util.GetVersion(),
		Env:     EnvStaging,
		Uptime:  uptime.String(),
	}
}
