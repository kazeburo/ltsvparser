# ltsvparser

LTSV (Labeled Tab-separated Values) parser for Go language.

ltsvparser provides bare API for performance. it's able to parse Millions LTSV record in a second.

This code is used in mackerel-plugin-axslog.

## Reference

### Each

```
func Each(d []byte, cb func(int, []byte) error, keys ...[]byte) error
```

This API refers EachKey of https://github.com/buger/jsonparser.

`Each` parses `[]byte` payload and calls callback function when key is found.

```
func main() {
	data := `
time:05/Feb/2013:15:34:47 +0000	host:192.168.50.1	req:GET / HTTP/1.1	status:200	reqtime:0.030
time:05/Feb/2013:15:35:15 +0000	host:192.168.50.1	req:GET /foo HTTP/1.1   status:200	reqtime:0.050
time:05/Feb/2013:15:35:54 +0000	host:192.168.50.1	req:GET /bar HTTP/1.1   status:404	reqtime:0.110
`
	b := bytes.NewBufferString(data)
	var statusOK = 0
	var statusNotOK = 0
	var totalReqTime = 0.0
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		err := ltsvparser.Each(
			scanner.Bytes(),
			func(idx int, v []byte) error {
				switch idx {
				case 0:
					// status
					if bytes.Equal(v, []byte("200")) {
						statusOK++
					} else {
						statusNotOK++
					}
				case 1:
					// reqtime
                    // Also can use jsonparser.ParseFloat
					rt, _ := strconv.ParseFloat(string(v)), 64)
					totalReqTime = totalReqTime + rt
				}
				return nil
			},
			[]byte("status"),
			[]byte("reqtime"),
		)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("OK:%d NotOK:%d TotalReqTime:%f\n", statusOK, statusNotOK, totalReqTime)
}
```

## Benchmarking

Benchmark codes https://gist.github.com/kazeburo/204efec4fab4a781f887ffa3d08a69c1

Parse 100k lines of LTSV

```
goos: darwin
goarch: amd64
pkg: github.com/kazeburo/go-ltsvparser-bench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkLtsv-8                        1        3005857179 ns/op        554412888 B/op  17100068 allocs/op
BenchmarkGoLtsv-8                      1        1020839457 ns/op        619206816 B/op   8100041 allocs/op
BenchmarkLtsvParser-8                  5         212723109 ns/op            4216 B/op          4 allocs/op
PASS
ok      github.com/kazeburo/go-ltsvparser-bench 6.474s
```

## Link

http://ltsv.org/

https://github.com/najeira/ltsv LTSV (Labeled Tab-separated Values) reader/writer for Go language.

https://github.com/Songmu/go-ltsv LTSV parser and encoder for Go with reflection

