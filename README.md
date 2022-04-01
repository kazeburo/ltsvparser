# ltsvparser

LTSV (Labeled Tab-separated Values) parser for Go language.

ltsvparser provides bare API for performance. ltsvparser can parse Millions LTSV record in a second.

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
                    // can use jsonparser.ParseFloat
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

## Link

http://ltsv.org/

https://github.com/najeira/ltsv LTSV (Labeled Tab-separated Values) reader/writer for Go language.

https://github.com/Songmu/go-ltsv LTSV parser and encoder for Go with reflection

