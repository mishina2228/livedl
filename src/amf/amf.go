package amf

import (
	"./amf0"
	"./amf_t"
	"bytes"
	"io"
)

func SwitchToAmf3() amf_t.SwitchToAmf3 {
	return amf_t.SwitchToAmf3{}
}

func EncodeAmf0(data []interface{}, asEcmaArray bool) ([]byte, error) {
	return amf0.Encode(data, asEcmaArray)
}

// paddingHint: zero padded before AMF data
func DecodeAmf0(data []byte, paddingHint ...bool) (res []interface{}, err error) {
	rdr := bytes.NewReader(data)
	var seek1 bool
	for _, h := range paddingHint {
		if h {
			seek1 = true
			break
		}
	}
	if seek1 {
		rdr.Seek(1, io.SeekStart)
	}
	res, err = amf0.DecodeAll(rdr)
	if err != nil {
		if seek1 {
			// retry
			rdr.Seek(0, io.SeekStart)
			res, err = amf0.DecodeAll(rdr)
		}
	}

	return
}
