package config

import (
	qrsDB "db/querys"
)

type IPGuardeConfig struct {
	Interval int // seconds: 1 int = 1s.
	RPSLimit int // request per seconds limit
	QRS      *qrsDB.QuerysDB
}
