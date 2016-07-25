
all:
	go-bindata -o geolite-db.go -pkg geolite ./GeoLite2-City.mmdb.bz2
