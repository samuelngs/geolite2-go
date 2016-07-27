
all:
	go-bindata -o geolite-db.go -pkg geolite ./GeoLite2-City.mmdb.bz2

test:
	go run test/main.go

run-test:
	bi
