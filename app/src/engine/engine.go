package engine

import (
	qrsDB "db/querys"
	cfg "firesimple/config"
	svc "firesimple/services"
	"sync"
)

type FireSPEngine struct {
	config cfg.FireSPConfig
	querys *qrsDB.QuerysDB
}

func New(config cfg.FireSPConfig, querys *qrsDB.QuerysDB) *FireSPEngine {
	return &FireSPEngine{
		config: config,
		querys: querys,
	}
}

func (fire *FireSPEngine) Start() *svc.FireSPServices {

	return &svc.FireSPServices{
		ServiceIPGuarde: fire.buildServiceIPGuarde(),
	}
}

func (fire *FireSPEngine) buildServiceIPGuarde() *svc.ServiceIPGuarde {
	c, _ := fire.querys.GetAllClientsIP()
	var clients sync.Map
	for _, v := range c {
		clients.Store(v, 0)
	}
	return svc.NewServiceIPGuarde(fire.config.IPGuarde, &clients)
}
