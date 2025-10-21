package services

import (
	qrsDB "db/querys"
	cfg "firesimple/config"
	"sync"
	"time"
)

type ServiceIPGuarde struct {
	clients  *sync.Map
	interval int
	rpsLimit int
	qrs      *qrsDB.QuerysDB
}

func NewServiceIPGuarde(config cfg.IPGuardeConfig, clients *sync.Map) *ServiceIPGuarde {
	service := &ServiceIPGuarde{
		clients:  clients,
		interval: config.Interval,
		rpsLimit: config.RPSLimit,
		qrs:      config.QRS,
	}
	go service.analyzeIPTraffic()
	return service
}

func (s *ServiceIPGuarde) analyzeIPTraffic() {
	for {
		duration := time.Duration(s.interval) * time.Second
		timer := time.NewTimer(duration)
		<-timer.C
		s.clients.Range(func(k, v any) bool {
			rps := float32(v.(int)) / 60

			if rps > float32(s.rpsLimit) {
				go s.qrs.UpdateClientSecure(k.(string), false)
				s.clients.Delete(k)
				return true
			}

			s.clients.Store(k, 0)
			return true
		})
	}
}
func (s *ServiceIPGuarde) IncrementIPCounter(clientIP string) {
	val, ok := s.clients.Load(clientIP)
	if !ok {
		s.clients.Store(clientIP, 1)
		return
	}
	s.clients.Store(clientIP, val.(int)+1)
}

func (s *ServiceIPGuarde) AuthClientIP(clientIP string) bool {
	isSecure, _ := s.qrs.GetStateSecureClient(clientIP)
	return isSecure
}
