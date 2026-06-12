package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <domain>")
		return
	}

	domain := os.Args[1]

	// A records (IPv4)
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("A records for %s:\n", domain)
	for _, ip := range ips {
		fmt.Printf("  %s\n", ip.String())
	}

	// MX records
	mxs, err := net.LookupMX(domain)
	if err == nil && len(mxs) > 0 {
		fmt.Printf("\nMX records for %s:\n", domain)
		for _, mx := range mxs {
			fmt.Printf("  %d %s\n", mx.Pref, mx.Host)
		}
	}

	// CNAME
	cname, err := net.LookupCNAME(domain)
	if err == nil {
		fmt.Printf("\nCNAME for %s: %s\n", domain, cname)
	}

	// TXT records
	txts, err := net.LookupTXT(domain)
	if err == nil && len(txts) > 0 {
		fmt.Printf("\nTXT records for %s:\n", domain)
		for _, txt := range txts {
			fmt.Printf("  %s\n", txt)
		}
	}
}
