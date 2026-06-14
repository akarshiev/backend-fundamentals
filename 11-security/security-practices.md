# Amaliy xavfsizlik (Security Practices)

Kunlik ishda qo'llaniladigan xavfsizlik usullari.

---

## HTTPS tekshirish

```bash
curl -I https://google.com
```

---

## Certificate tekshirish

```bash
openssl s_client -connect google.com:443
```

---

## ARP table

```bash
arp -a
```

---

## DNS

```bash
dig google.com
```

---

## TCP ulanishlar

```bash
ss -tan
```

---

## Portlar

```bash
ss -tulpn
```

---

## Traffic Capture

```bash
sudo tcpdump -i any
```

yoki Wireshark

---

## SSL/TLS tekshirish

```bash
# TLS versiyasini tekshirish
openssl s_client -connect google.com:443

# Sertifikat muddatini tekshirish
openssl s_client -connect google.com:443 </dev/null 2>/dev/null | openssl x509 -noout -dates
```

---

## Firewall sozlash

```bash
# Ulanishlarni ko'rish
sudo iptables -L -n

# Qoida qo'shish
sudo iptables -A INPUT -p tcp --dport 80 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 443 -j ACCEPT
sudo iptables -A INPUT -p tcp -j DROP
```

---

## Xulosa

- HTTPS, SSL/TLS ni tekshiring
- ARP, DNS, TCP ulanishlarini kuzating
- Traffic capture vositalarini ishlating
