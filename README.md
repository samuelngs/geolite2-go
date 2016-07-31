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

## Documentation

`go doc` format documentation for this project can be viewed online without installing the package by using the GoDoc page at: https://godoc.org/github.com/samuelngs/geolite2-go

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/samuelngs/geolite2-go/issues)
- Fix bugs and [submit pull requests](https://github.com/samuelngs/geolite2-go/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features

## License ##

This project is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.

```
The MIT License (MIT)

Copyright (c) 2016 Samuel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
