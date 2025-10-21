module firesimple/engine

go 1.24.7

replace firesimple/config => ../config

replace firesimple/services => ../services

replace db/querys => ../../db/querys

require firesimple/config v0.0.0-00010101000000-000000000000

require (
	db/querys v0.0.0-00010101000000-000000000000
	firesimple/services v0.0.0-00010101000000-000000000000
)
