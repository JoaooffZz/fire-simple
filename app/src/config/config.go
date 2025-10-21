package config

import (
	qrsDB "db/querys"
	"os"
	"strconv"
)

type FireSPConfig struct {
	IPGuarde IPGuardeConfig
}

func Default(querys *qrsDB.QuerysDB) FireSPConfig {
	return FireSPConfig{
		IPGuarde: buildIPGuarde(querys),
	}
}

func buildIPGuarde(qrs *qrsDB.QuerysDB) IPGuardeConfig {
	var interval int
	var rpsLimit int
	strInterval := os.Getenv("FIRESIMPLE_CONFIG_IPGUARDE_INTERVAL")
	if strInterval == "" {
		interval = 60
	} else {
		n, err := strconv.Atoi(strInterval)
		if err != nil {
			n = 60
		}
		interval = n
	}

	strRpsLimit := os.Getenv("FIRESIMPLE_CONFIG_IPGUARDE_REQUEST_SECONDS_LIMIT")
	if strRpsLimit == "" {
		rpsLimit = 15
	} else {
		n, err := strconv.Atoi(strInterval)
		if err != nil {
			n = 15
		}
		rpsLimit = n
	}

	return IPGuardeConfig{
		Interval: interval,
		RPSLimit: rpsLimit,
		QRS:      qrs,
	}
}
