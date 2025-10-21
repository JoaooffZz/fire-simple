module server

go 1.24.7

replace db/conn => ../db/conn

replace db/querys => ../db/querys

replace firesimple/engine => ../src/engine

replace firesimple/services => ../src/services

replace firesimple/config => ../src/config

require firesimple/services v0.0.0-00010101000000-000000000000

require (
	db/querys v0.0.0-00010101000000-000000000000 // indirect
	firesimple/config v0.0.0-00010101000000-000000000000 // indirect
)
