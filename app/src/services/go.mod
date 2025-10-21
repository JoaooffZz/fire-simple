module firesimple/services

go 1.24.7

replace firesimple/config => ../config

replace db/querys => ../../db/querys

require (
	db/querys v0.0.0-00010101000000-000000000000
	firesimple/config v0.0.0-00010101000000-000000000000
)
