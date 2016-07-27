package geolite

import "testing"

// TestBytes func
func TestBytes(t *testing.T) {
	switch b, err := Bytes("GeoLite2-City.mmdb"); {
	case err != nil:
		t.Error(err)
	default:
		t.Logf("database size: %d", len(b))
	}
}
