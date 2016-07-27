# GeoLite2 embedding binary

This product includes GeoLite2 data created by MaxMind, available from http://www.maxmind.com.

```

import (
  "log"
  "github.com/samuelngs/geolite2-go"
)

func init() {
  db, err := geolite.Bytes("GeoLite2-City.mmdb")
  if err != nil {
    log.Fatal(err)
  }
  ...
}

```
