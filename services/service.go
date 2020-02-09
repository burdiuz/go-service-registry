package services

import (
	"time"
)

type Service struct {
	Name      string
	Version   string
	IPAddress string
	Port      string
	createdAt time.Time
}

func ServiceNew(name, version, ipaddr, port string) *Service {
	return &Service{name, version, ipaddr, port, time.Now()}
}
