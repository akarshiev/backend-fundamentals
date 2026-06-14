package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run dns-client.go <domain>")
		fmt.Println("Example: go run dns-client.go google.com")
		return
	}

	domain := os.Args[1]

	fmt.Printf("DNS lookup for %s:\n", domain)
	fmt.Println(strings.Repeat("-", 40))

	// A records (IPv4)
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("A records (IPv4):")
	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Printf("  %s\n", ip.String())
		}
	}

	// AAAA records (IPv6)
	fmt.Println("\nAAAA records (IPv6):")
	for _, ip := range ips {
		if ip.To4() == nil {
			fmt.Printf("  %s\n", ip.String())
		}
	}

	// MX records
	mxs, err := net.LookupMX(domain)
	if err == nil && len(mxs) > 0 {
		fmt.Println("\nMX records:")
		for _, mx := range mxs {
			fmt.Printf("  %d %s\n", mx.Pref, mx.Host)
		}
	}

	// CNAME
	cname, err := net.LookupCNAME(domain)
	if err == nil {
		fmt.Printf("\nCNAME: %s\n", cname)
	}

	// TXT records
	txts, err := net.LookupTXT(domain)
	if err == nil && len(txts) > 0 {
		fmt.Println("\nTXT records:")
		for _, txt := range txts {
			fmt.Printf("  %s\n", txt)
		}
	}

	// NS records
	nss, err := net.LookupNS(domain)
	if err == nil && len(nss) > 0 {
		fmt.Println("\nNS records:")
		for _, ns := range nss {
			fmt.Printf("  %s\n", ns.Host)
		}
	}

	// PTR record (reverse lookup)
	fmt.Println("\nReverse lookup:")
	for _, ip := range ips {
		names, err := net.LookupAddr(ip.String())
		if err == nil && len(names) > 0 {
			for _, name := range names {
				fmt.Printf("  %s -> %s\n", ip.String(), name)
			}
		}
	}
}
