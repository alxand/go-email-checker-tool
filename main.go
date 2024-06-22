package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMAR bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMx = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=dmarc1") {
			hasDMAR = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v %v, %v", domain, hasMx, hasSPF, spfRecord, hasDMAR, dmarcRecord)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, hasDMAR, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
