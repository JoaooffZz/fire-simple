package server

import (
	svc "firesimple/services"
)

func handlerIPGuarde(services *svc.FireSPServices, typeService, clientIP string, write func(b []byte) (n int, err error)) {
	switch typeService {
	case "auth.client-ip":
	default:
	}
}
