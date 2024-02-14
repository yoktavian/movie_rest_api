.PHONY: m-create

m-create:
	migrate create -ext sql -dir db/migration -seq $(name)
m-up:
	migrate -path db/migration/ -database "mysql://root:@tcp(127.0.0.1:3306)/playground-movie" -verbose up
m-down:
	migrate -path db/migration/ -database "mysql://root:@tcp(127.0.0.1:3306)/playground-movie" -verbose down 1

name ?= default_argument
