
all:
	binary -dir ./GeoLite2-City.mmdb -pkg geolite -max 41943040

test:
	go test -v

