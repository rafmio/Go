module logtracker

go 1.22.2

replace parser => ./parser

replace dbhandler => ./dbhandler

require (
	github.com/lib/pq v1.10.9 // indirect
	parser v0.0.0-00010101000000-000000000000 // indirect
)
