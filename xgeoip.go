package xgeoip

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/netip"
	"os"
)

const defaultV4DBFileName = "GeoLite2-ASN-Blocks-IPv4.csv"

func Init(dbFileName string) error {
	if dbFileName == "" {
		fmt.Println("db file name was not provided. Default will be used", defaultV4DBFileName)
	}
	dbFileName = defaultV4DBFileName
	if err := parseCsvDB(dbFileName); err != nil {
		return err
	}
	if st.isEmpty() {
		log.Fatalln("empty maxmind DB")
	}
	return nil
}

func parseCsvDB(dbFileName string) error {
	f, err := os.Open(dbFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		save(rec)
	}
	return nil
}

func save(rec []string) {
	if len(rec) != 3 {
		return
	}
	prefix, err := netip.ParsePrefix(rec[0])
	if err != nil {
		return
	}
	st.Add(prefix, rec[1], rec[2])
}

func GetByAddr(ip string) *Info {
	return st.GetByAddr(ip)
}

func GetByASN(asn string) []*Info {
	return st.GetByASN(asn)
}

func GetDB() map[string]*Info {
	return st.data
}
