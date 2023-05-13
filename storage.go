package xgeoip

import (
	"net/netip"
	"sync"
	"time"
)

type storage struct {
	sync.Mutex
	lastUpdate time.Time
	data       map[string]*Info
}

var st = &storage{data: make(map[string]*Info, 1024)}

func (s *storage) isEmpty() bool {
	if len(s.data) > 0 {
		return false
	}
	return true
}

func (s *storage) Add(prefix netip.Prefix, asn, org string) {
	if len(prefix.String()) == 0 || len(asn) == 0 || len(org) == 0 {
		return
	}
	s.Lock()
	defer s.Unlock()
	if s.isEmpty() {
		s.data = make(map[string]*Info, 1024)
	}
	if _, ok := s.data[asn]; !ok {
		s.data[asn] = &Info{
			Net: []netip.Prefix{prefix},
			ASN: asn,
			Org: org,
		}
	} else {
		s.data[asn].Net = append(s.data[asn].Net, prefix)
	}
	s.lastUpdate = time.Now()
}

func (s *storage) GetByAddr(ip string) *Info {
	if s.isEmpty() {
		return nil
	}
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return nil
	}
	for _, data := range s.data {
		for _, prefix := range data.Net {
			if prefix.Contains(addr) {
				return data
			}
		}
	}
	return nil
}

func (s *storage) GetByASN(asn string) []*Info {
	if s.isEmpty() || len(asn) == 0 {
		return nil
	}
	ii := make([]*Info, 0)
	for _, data := range s.data {
		if data.ASN == asn {
			ii = append(ii, data)
		}
	}
	if len(ii) > 0 {
		return ii
	}
	return nil
}
