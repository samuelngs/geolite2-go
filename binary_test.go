package geolite

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

// TestBytes func
func TestBytes(t *testing.T) {
	switch b, err := Bytes("GeoLite2-City.mmdb"); {
	case err != nil:
		t.Error(err)
	default:
		hash := md5.New()
		hash.Write(b[:])
		t.Logf("database size: %v, md5 hash: %v", len(b), hex.EncodeToString(hash.Sum(nil)))
	}
}
