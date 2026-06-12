import socket
import sys

def dns_lookup(domain):
    print(f"A records for {domain}:")
    ips = socket.getaddrinfo(domain, 80, socket.AF_INET)
    for ip in ips:
        print(f"  {ip[4][0]}")

    try:
        ips_v6 = socket.getaddrinfo(domain, 80, socket.AF_INET6)
        if ips_v6:
            print(f"\nAAAA records for {domain}:")
            for ip in ips_v6:
                print(f"  {ip[4][0]}")
    except socket.gaierror:
        pass

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <domain>")
        sys.exit(1)

    dns_lookup(sys.argv[1])
