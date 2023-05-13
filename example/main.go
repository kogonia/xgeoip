package main

import (
	"flag"
	"fmt"
	"github.com/kogonia/xgeoip"
	"log"
)

var (
	addr = flag.String("ip", "", "get info for specific IP address")
	asn  = flag.String("asn", "", "get info for specific ASN")
)

func main() {
	flag.Parse()

	if err := xgeoip.Init(""); err != nil {
		log.Fatal(err)
	}

	switch {
	case len(*addr) > 0:
		addrInfo := xgeoip.GetByAddr(*addr)
		fmt.Println(addrInfo)
	case len(*asn) > 0:
		asnInfo := xgeoip.GetByASN(*asn)
		fmt.Println(asnInfo)
	default:
		fmt.Println("Empty query. Usage:")
		flag.PrintDefaults()
	}

}
