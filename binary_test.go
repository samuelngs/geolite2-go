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
		sum := md5.New()
		sum.Write(b[:])
		hash := hex.EncodeToString(sum.Sum(nil))
		expected := "b799f8c5fa32ab0b17a6ce9ee2f7e3ef"
		if hash != expected {
			t.Fatalf("database size %v KB, expected md5 %v, got %v", len(b), expected, hash)
		} else {
			t.Logf("database size: %v KB, md5 %v", len(b), hash)
		}
	}
}
