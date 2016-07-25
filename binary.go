package geolite

import (
	"bytes"
	"compress/gzip"
	"errors"
	"log"
)

var file = map[string][][]byte{
	"GeoLite2-City.mmdb": [][]byte{
		d71a359a4fb80c46da5fdf0cad2a8f0572c203eeb, da6921c8c7d2adb084f765602c6bebf2ee5c2d127, dd59086da7b889891ec359fd2d6a447431f6ad933,
	},
}

var size = map[string]int{
	"GeoLite2-City.mmdb": 65865934,
}

// Get returns file in bytes format
func Get(filename string) ([]byte, error) {
	var b bytes.Buffer
	data, ok := file[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	length := size[filename]
	for _, part := range data {
		b.Write(part)
	}
	res := make([]byte, length)
	r := bytes.NewReader(b.Bytes())
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	if _, err := gz.Read(res); err != nil {
		return nil, err
	}
	return res, nil
}

// MustGet to retrieve data, panic if fail
func MustGet(filename string) []byte {
	b, err := Get(filename)
	if err != nil {
		log.Panic(err)
	}
	return b
}
