package ltsvparser

import (
	"bytes"
)

var tab = []byte("\t")
var col = []byte(":")
var hif = []byte("-")
var null = []byte("")

type Canceler struct{}

func (e *Canceler) Error() string {
	return ""
}

// Stop parser without errors
var Cancel = &Canceler{}

// Extract multiple keys from LTSV
func Each(d []byte, cb func(int, []byte) error, keys ...[]byte) error {
	p1 := 0
	dlen := len(d)
	for {
		if dlen <= p1 {
			break
		}
		p2 := bytes.Index(d[p1:], tab)
		if p2 == 0 { // first byte is tab
			p1 += p2 + 1
			continue
		}
		if p2 < 0 { // could not find tab
			p2 = dlen - p1
		}
		p3 := bytes.Index(d[p1:p1+p2], col)
		if p3 < 0 { // could not find :
			p3 = p2
		}
		for i := range keys {
			if bytes.Equal(d[p1:p1+p3], keys[i]) {
				// value as '-'
				var cbErr error
				if p3 == p2 || bytes.Equal(d[p1+p3+1:p1+p2], hif) {
					cbErr = cb(i, null)
				} else {
					cbErr = cb(i, d[p1+p3+1:p1+p2])
				}
				if cbErr != nil {
					if _, ok := cbErr.(*Canceler); ok {
						return nil
					}
					return cbErr
				}
			}
		}
		p1 += p2 + 1
	}
	return nil
}
