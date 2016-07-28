
all:
	binary -dir ./GeoLite2-City.mmdb -pkg geolite -max 20971520

test:
	go test -v

