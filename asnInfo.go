package xgeoip

import (
	"bytes"
	"encoding/json"
	"net/netip"
	"strings"
)

type Info struct {
	Net []netip.Prefix
	ASN string
	Org string
}

func (info *Info) String() string {
	js, _ := json.MarshalIndent(info, "", "\t")
	return string(js)
}

func (info *Info) Json() string {
	return strings.TrimSpace(string(info.bytes()))
}

func (info *Info) bytes() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 64))
	_ = json.NewEncoder(b).Encode(info)
	return b.Bytes()
}
